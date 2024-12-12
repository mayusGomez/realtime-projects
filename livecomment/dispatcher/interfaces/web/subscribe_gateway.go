package web

import (
	"github.com/gin-gonic/gin"
	"livecomments/dispatcher/domain"
	"log"
	"net/http"
)

type GatewaySubscription struct {
	IsSubscription *bool  `json:"is_subscription"`
	Queue          string `json:"queue" binding:"required"`
	Video          string `json:"video" binding:"required"`
}

type SubscribeGateway struct {
	subscriptionCmd domain.SubscriptionCmd
}

func NewSubscribeGateway(subscriptionCmd domain.SubscriptionCmd) *SubscribeGateway {
	return &SubscribeGateway{
		subscriptionCmd: subscriptionCmd,
	}
}

func (gw *SubscribeGateway) Handle(c *gin.Context) {
	var req GatewaySubscription

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Received GatewaySubscription %+v", req)

	var err error
	if req.IsSubscription != nil && !*req.IsSubscription {
		err = gw.subscriptionCmd.Unsubscribe(req.Queue, req.Video)
	} else {
		err = gw.subscriptionCmd.Subscribe(req.Queue, req.Video)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
