package commentgenerator

import (
	"bytes"
	"encoding/json"
	faker "github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

func CommentGenerator(dispatcherURL, video string) {
	for {
		time.Sleep(1 * time.Second)

		payload := map[string]interface{}{
			"video":         video,
			"connection_id": uuid.New().String(),
			"comment":       faker.Comment(),
		}

		jsonData, err := json.Marshal(payload)
		if err != nil {
			log.Printf("error encoding JSON: %v\n", err)
			return
		}

		resp, err := http.Post(dispatcherURL, "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Printf("error sending POST request to %s: %v\n", dispatcherURL, err)
			continue
		}
		resp.Body.Close()
	}
}
