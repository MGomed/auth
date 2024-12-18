package app

import (
	"context"
	"errors"
	"flag"
	"io"
	"log"
	"net"
	"net/http"
	"sync"

	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	fs "github.com/rakyll/statik/fs"
	cors "github.com/rs/cors"
	grpc "google.golang.org/grpc"
	credentials "google.golang.org/grpc/credentials"
	insecure "google.golang.org/grpc/credentials/insecure"
	reflection "google.golang.org/grpc/reflection"

	consts "github.com/MGomed/auth/consts"
	interceptors "github.com/MGomed/auth/internal/api/interceptors"
	"github.com/MGomed/auth/internal/api/metrics"
	env_config "github.com/MGomed/auth/internal/config/env"
	access_api "github.com/MGomed/auth/pkg/access_api"
	auth_api "github.com/MGomed/auth/pkg/auth_api"
	user_api "github.com/MGomed/auth/pkg/user_api"
	closer "github.com/MGomed/common/closer"

	// Needed to get static files
	_ "github.com/MGomed/auth/pkg/statik"
)

var configPath string

// App represents object for starting grpc server
type App struct {
	ctx              context.Context
	serviceProvider  *serviceProvider
	grpcServer       *grpc.Server
	httpServer       *http.Server
	swaggerServer    *http.Server
	prometheusServer *http.Server
}

// NewApp is App struct constructor
func NewApp(ctx context.Context) (*App, error) {
	flag.StringVar(&configPath, "config-path", "build/.env", "path to config file")
	flag.Parse()

	app := &App{
		ctx: ctx,
	}

	if err := app.initDeps(ctx); err != nil {
		return nil, err
	}

	return app, nil
}

// Run starts grpc server
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		err := a.runGRPCServer()
		if err != nil {
			log.Fatalf("failed to run GRPC server: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := a.runHTTPServer()
		if err != nil {
			log.Fatalf("failed to run HTTP server: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := a.runSwaggerServer()
		if err != nil {
			log.Fatalf("failed to run Swagger server: %v", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		err := a.runPrometheusServer()
		if err != nil {
			log.Fatalf("failed to run Prometheus server: %v", err)
		}
	}()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()

	// 	consumerCreate := a.serviceProvider.Consumer()
	// 	closer.Add(consumerCreate.Close)

	// 	err := consumerCreate.Consume(a.ctx, consts.CreateTopic, kafka_consumer.Handler(
	// 		func(_ context.Context, msg *sarama.ConsumerMessage) error {
	// 			logger := a.serviceProvider.Logger()

	// 			logger.Printf("MESSAGE FROM KAFKA: >>> User created:\n%v\n", msg.Value)

	// 			return nil
	// 		}),
	// 	)

	// 	if err != nil {
	// 		a.serviceProvider.Logger().Println("consumer error: ", err)
	// 	}
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()

	// 	consumerDelete := a.serviceProvider.Consumer()
	// 	closer.Add(consumerDelete.Close)

	// 	err := consumerDelete.Consume(a.ctx, consts.DeleteTopic, kafka_consumer.Handler(
	// 		func(_ context.Context, msg *sarama.ConsumerMessage) error {
	// 			logger := a.serviceProvider.Logger()

	// 			logger.Printf("MESSAGE FROM KAFKA: >>> User deleted:\n%v\n", msg.Value)

	// 			return nil
	// 		}),
	// 	)

	// 	if err != nil {
	// 		a.serviceProvider.Logger().Println("consumer error: ", err)
	// 	}
	// }()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHTTPServer,
		a.initSwaggerServer,
		a.initPrometheusServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	if err := env_config.Load(configPath); err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	creds, err := credentials.NewServerTLSFromFile(consts.ServiceCertFilePath, consts.ServiceCertKeyFilePath)
	if err != nil {
		return err
	}

	a.grpcServer = grpc.NewServer(
		grpc.Creds(creds),
		grpc.ChainUnaryInterceptor(
			interceptors.ValidateInterceptor,
			interceptors.MetricsInterceptor,
		),
	)

	reflection.Register(a.grpcServer)

	user_api.RegisterUserAPIServer(a.grpcServer, a.serviceProvider.UserAPI(ctx))
	auth_api.RegisterAuthAPIServer(a.grpcServer, a.serviceProvider.AuthAPI(ctx))
	access_api.RegisterAccessAPIServer(a.grpcServer, a.serviceProvider.AccessAPI())

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := user_api.RegisterUserAPIHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts)
	if err != nil {
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:              a.serviceProvider.HTTPConfig().Address(),
		Handler:           corsMiddleware.Handler(mux),
		ReadHeaderTimeout: consts.ReadHeaderTimeout,
	}

	return nil
}

func (a *App) initSwaggerServer(_ context.Context) error {
	statikFs, err := fs.New()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.StripPrefix("/", http.FileServer(statikFs)))
	mux.HandleFunc(consts.SwaggerPath, serveSwaggerFile(consts.SwaggerPath))

	a.swaggerServer = &http.Server{
		Addr:              a.serviceProvider.SwaggerConfig().Address(),
		Handler:           mux,
		ReadHeaderTimeout: consts.ReadHeaderTimeout,
	}

	return nil
}

func (a *App) initPrometheusServer(ctx context.Context) error {
	if err := metrics.Init(ctx); err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheusServer = &http.Server{
		Addr:              a.serviceProvider.PrometheusConfig().Address(),
		Handler:           mux,
		ReadHeaderTimeout: consts.ReadHeaderTimeout,
	}

	return nil
}

func (a *App) runGRPCServer() error {
	lis, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}
	closer.Add(func() error {
		a.grpcServer.Stop()

		return nil
	})

	return a.grpcServer.Serve(lis)
}

func (a *App) runHTTPServer() error {
	closer.Add(a.httpServer.Close)
	if err := a.httpServer.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}

func (a *App) runSwaggerServer() error {
	closer.Add(a.swaggerServer.Close)
	if err := a.swaggerServer.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}

func (a *App) runPrometheusServer() error {
	closer.Add(a.prometheusServer.Close)
	if err := a.prometheusServer.ListenAndServe(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
	}

	return nil
}

func serveSwaggerFile(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		statikFs, err := fs.New()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		file, err := statikFs.Open(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer func() {
			_ = file.Close()
		}()

		content, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}
