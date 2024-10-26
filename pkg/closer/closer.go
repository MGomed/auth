package closer

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/MGomed/auth/pkg/logger"
	"golang.org/x/sync/errgroup"
)

var globalCloser = New(logger.InitLogger(), syscall.SIGTERM, syscall.SIGINT)

// Add adds `func() error` callback to the globalCloser
func Add(f ...func() error) {
	globalCloser.Add(f...)
}

// Wait ...
func Wait() {
	globalCloser.Wait()
}

// CloseAll ...
func CloseAll() {
	globalCloser.CloseAll()
}

// Closer ...
type Closer struct {
	log   *log.Logger
	mu    sync.Mutex
	once  sync.Once
	done  chan struct{}
	funcs []func() error
}

// New returns new Closer, if []os.Signal is specified Closer will automatically call CloseAll when one of signals is received from OS
func New(log *log.Logger, sig ...os.Signal) *Closer {
	c := &Closer{
		log:  log,
		done: make(chan struct{}),
	}
	if len(sig) > 0 {
		go func() {
			ch := make(chan os.Signal, 1)
			signal.Notify(ch, sig...)
			<-ch
			signal.Stop(ch)
			c.CloseAll()
		}()
	}
	return c
}

// Add func to closer
func (c *Closer) Add(f ...func() error) {
	c.mu.Lock()
	c.funcs = append(c.funcs, f...)
	c.mu.Unlock()
}

// Wait blocks until all closer functions are done
func (c *Closer) Wait() {
	<-c.done
	fmt.Println("HELLLLLLO")
}

// CloseAll calls all closer functions
func (c *Closer) CloseAll() {
	c.once.Do(func() {
		defer close(c.done)

		c.mu.Lock()
		funcs := c.funcs
		c.funcs = nil
		c.mu.Unlock()

		// call all Closer funcs async
		g, _ := errgroup.WithContext(context.Background())
		for _, f := range funcs {
			g.Go(f)
		}

		if err := g.Wait(); err != nil {
			fmt.Println("AJJAJAAJ")
			log.Printf("error returned from Closer: %v\n", err) 
		}
	})
}
