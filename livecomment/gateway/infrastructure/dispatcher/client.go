package dispatcher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"livecomments/pkg/httpstandarclient"
	"log"
	"net/http"
)

const subscriptionPath = "/gateway/subscription"

type Client struct {
	url    string
	client *http.Client
}

func NewDispatcher(url string) *Client {
	return &Client{
		url:    url,
		client: httpstandarclient.DefaultHTTPConfig(3, 3),
	}
}

func (dispatcher *Client) Subscribe(video, queue string) error {
	payload := map[string]interface{}{
		"video":           video,
		"queue":           queue,
		"is_subscription": true,
	}

	return dispatcher.callServer(subscriptionPath, payload)
}

func (dispatcher *Client) Unsubscribe(video, queue string) error {
	payload := map[string]interface{}{
		"video":           video,
		"queue":           queue,
		"is_subscription": false,
	}

	return dispatcher.callServer(subscriptionPath, payload)
}

func (dispatcher *Client) callServer(path string, payload map[string]interface{}) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal payload: %v", err)
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s%s", dispatcher.url, path), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("failed to create request: %v", err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	resp, err := dispatcher.client.Do(req)
	if err != nil {
		log.Printf("failed to send request: %v", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body: %v", err)
		return err
	}

	log.Printf("Response status: %s\n", resp.Status)
	log.Printf("Response body: %s\n", string(body))

	return nil
}
