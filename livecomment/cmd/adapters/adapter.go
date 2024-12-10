package adapters

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/vardius/shutdown"
)

const (
	defaultShutdownTimeout = 40 * time.Second
)

type Adapter interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}

type App struct {
	adapters        []Adapter
	shutdownTimeout time.Duration
}

func NewAppAdapters() *App {
	return &App{
		shutdownTimeout: defaultShutdownTimeout, // Default shutdown timeout
	}
}

func (app *App) AddAdapters(adapters ...Adapter) {
	app.adapters = append(app.adapters, adapters...)
}

func (app *App) Run() {
	wg := &sync.WaitGroup{}

	for _, adapter := range app.adapters {
		wg.Add(1)

		go func(adapter Adapter) {
			defer wg.Done()

			if err := adapter.Start(context.Background()); err != nil {
				log.Printf("failed to start adapter: %v", err)
			}
		}(adapter)
	}

	shutdown.GracefulStop(func() { app.Stop(context.Background()) })
	wg.Wait()
}

func (app *App) Stop(ctx context.Context) {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, app.shutdownTimeout)
	defer cancel()

	errCh := make(chan error, len(app.adapters))

	for _, adapter := range app.adapters {
		go func(adapter Adapter) {
			errCh <- adapter.Stop(ctxWithTimeout)
		}(adapter)
	}

	for i := 0; i < len(app.adapters); i++ {
		if err := <-errCh; err != nil {
			go func(err error) {
				log.Printf("failed to stop adapter: %v", err)
				os.Exit(1)
			}(err)

			return
		}
	}
}
