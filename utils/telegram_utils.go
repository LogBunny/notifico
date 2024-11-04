package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"logbun/config"
	"logbun/schemas"
	"net/http"
)

func SendNotificationMessage(chatID, message string) error {
	telegramMessage := schemas.TelegramMessageSchema{
		ChatID: chatID,
		Text:   message,
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", config.TELEGRAM_BOT_TOKEN)

	jsonData, err := json.Marshal(telegramMessage)
	if err != nil {
		log.Println(err)
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println(err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Println(fmt.Errorf("failed to send message to Telegram: %s", resp.Status))
		return fmt.Errorf("failed to send message to Telegram: %s", resp.Status)
	}
	return nil
}
