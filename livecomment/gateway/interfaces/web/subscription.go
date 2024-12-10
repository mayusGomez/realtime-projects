package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type SubscriptionHandler struct {
	SubscriptionCommand func(video, connectionId string) (chan string, error)
	UnsubscribeCommand  func(video, connectionId string)
}

func NewSubscriptionHandler(subsCommand func(video, connectionId string) (chan string, error), unsubCommand func(video, connectionId string)) *SubscriptionHandler {
	return &SubscriptionHandler{
		SubscriptionCommand: subsCommand,
		UnsubscribeCommand:  unsubCommand,
	}
}

func (s *SubscriptionHandler) Handle(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	c.Writer.Flush()

	ctx := c.Request.Context()

	video := c.Query("video")
	if video == "" {
		log.Println("You need to specify video")
		c.JSON(http.StatusBadRequest, gin.H{"error": "you need to specify video"})
		return
	}

	connection := uuid.New().String()

	ch, err := s.SubscriptionCommand(video, connection)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for {
		select {
		case <-ctx.Done(): // Web connection closed
			log.Println("Disconnected Client")
			s.UnsubscribeCommand(video, connection)
			return

		case msg := <-ch:
			log.Printf("Received message: %s, video: %s, connection: %s", msg, video, connection)
			_, err := fmt.Fprintf(c.Writer, "data: %s\n\n", msg)
			if err != nil {
				log.Println("Error writing to client", err)
				return
			}

			c.Writer.Flush()
		}
	}
}
