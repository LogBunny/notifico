package config

import "github.com/spf13/viper"

var (
	DB_URI                  = ""
	MIGRATE            bool = false
	TELEGRAM_BOT_TOKEN      = ""
)

func LoadCfg() {
	DB_URI = viper.GetString("DB_URL")
	MIGRATE = viper.GetBool("MIGRATE")
	TELEGRAM_BOT_TOKEN = viper.GetString("TELEGRAM_BOT_TOKEN")
}
