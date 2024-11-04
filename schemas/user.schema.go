package schemas

type UserSchema struct {
	UserName string `json:"user_name"`
	Email    string `json:"email"`
}

type UserUpdateSchema struct {
	UserName   string `json:"user_name"`
	Name       string `json:"name"`
	Email      string `gorm:"unique; not null" json:"email"`
	TelegramId string `json:"telegram_id"`
}

type UserEmailSchema struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}
