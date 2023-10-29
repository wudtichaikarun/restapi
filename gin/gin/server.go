package gin

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/wudtichaikarun/restapi"
	"github.com/wudtichaikarun/restapi/configs/configs"
)

type GinRestAPI struct {
	server *http.Server
	Config *configs.Configuration
	Router restapi.Router
}

func NewGinRestAPI(config *configs.Configuration, router *GinRouter) restapi.RestAPI {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port),
		Handler: router.engine,
	}
	return &GinRestAPI{Config: config, server: server}
}

// implements restapi.RestAPI
func (r *GinRestAPI) Start() error {
	fmt.Printf("Starting server on %s:%d", r.Config.Server.Host, r.Config.Server.Port)

	go func() {
		if err := r.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	if err := r.gracefulShutdown(); err != nil {
		return err
	}

	return nil
}

func (r *GinRestAPI) gracefulShutdown() error {
	//  Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	fmt.Println("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	fmt.Println("Stopping server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := r.server.Shutdown(ctx); err != nil {
		fmt.Printf("Server forced to shutdown:: %v\n", err)
		return err
	}
	fmt.Println("Server stopped.")
	return nil
}
