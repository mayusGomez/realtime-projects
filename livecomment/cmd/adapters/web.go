package adapters

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"livecomments/gateway"
	"log"
	"net/http"
)

type WebAdapter struct {
	serviceContainer *gateway.ServiceContainer
	port             string
	server           *http.Server
}

func NewWebAdapter(port string, serviceContainer *gateway.ServiceContainer) *WebAdapter {
	return &WebAdapter{
		serviceContainer: serviceContainer,
		port:             port,
	}
}

func (w *WebAdapter) Start(_ context.Context) error {
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/subscribe", w.serviceContainer.SubscriptionHandler)

	w.server = &http.Server{
		Addr:    fmt.Sprintf(":%s", w.port),
		Handler: router,
	}

	go func() {
		if err := w.server.ListenAndServe(); err != nil {
			log.Fatalf("Failed to run server: %v", err)
		}
	}()

	return nil
}

func (w *WebAdapter) Stop(_ context.Context) error {
	err := w.server.Shutdown(context.Background())
	if err != nil {
		return err
	}

	return nil
}
