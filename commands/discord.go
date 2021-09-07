package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func NotifyDiscord(message string) {
	type bodyWebhook struct {
		Content string `json:"content"`
	}
	body := &bodyWebhook{
		Content: message,
	}

	bodyToJson, _ := json.Marshal(&body)
	http.Post(
		os.Getenv("WEBHOOK"),
		"application/json", bytes.NewBuffer(bodyToJson))
	fmt.Println("Notify sended success")
}
