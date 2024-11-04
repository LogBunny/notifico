package schemas

type TelegramMessageSchema struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}
