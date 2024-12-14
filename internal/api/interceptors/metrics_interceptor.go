package interceptors

import (
	"context"
	"time"

	grpc "google.golang.org/grpc"

	api_metrics "github.com/MGomed/auth/internal/api/metrics"
)

// MetricsInterceptor grpc UnaryInterceptor interface implementation for metrics collection
func MetricsInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	api_metrics.IncRequestCounter()

	timeStart := time.Now()

	resp, err := handler(ctx, req)
	diffTime := time.Since(timeStart)

	if err != nil {
		api_metrics.IncResponseCounter("error", info.FullMethod)
		api_metrics.HistogramResponseTimeObserve("error", diffTime.Seconds())
	} else {
		api_metrics.IncResponseCounter("success", info.FullMethod)
		api_metrics.HistogramResponseTimeObserve("success", diffTime.Seconds())
	}

	return resp, nil
}
