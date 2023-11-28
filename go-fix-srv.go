package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const webhookURL = "" //discord webhook url

type DiscordMessage struct {
	Content string `json:"content"`
}

func sendDiscordMessage(message string) error {
	discordMessage := DiscordMessage{
		Content: message,
	}
	messageBytes, err := json.Marshal(discordMessage)
	if err != nil {
		return fmt.Errorf("error encoding JSON: %v", err)
	}
	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(messageBytes))
	if err != nil {
		return fmt.Errorf("error sending POST request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNoContent || resp.StatusCode == http.StatusOK {
		fmt.Println("Message sent to Discord successfully" + "," + "message:" + message)
		return nil
	}
	return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
}

func main() {
	var message string
	fmt.Print("enter message:")
	fmt.Scan(&message)
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := sendDiscordMessage(message)
			if err != nil {
				fmt.Printf("Error sending Discord message: %v\n", err)
			}
		}
	}
}
