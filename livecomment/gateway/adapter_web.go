package gateway

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type WebAdapter struct {
	subscriptHandler func(c *gin.Context)
	port             string
	server           *http.Server
}

func NewWebAdapter(port string, subscriptHandler func(c *gin.Context)) *WebAdapter {
	return &WebAdapter{
		subscriptHandler: subscriptHandler,
		port:             port,
	}
}

func (w *WebAdapter) Start(_ context.Context) error {
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/subscribe", w.subscriptHandler)

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
