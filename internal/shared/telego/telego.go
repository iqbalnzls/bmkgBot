package telego

import (
	"github.com/SakoDroid/telego"
	"github.com/SakoDroid/telego/configs"

	"github.com/bmkgBot/internal/shared/config"
)

func NewTelego(config *config.TelebotConfig) *telego.Bot {
	telegoCfg := &configs.BotConfigs{
		BotAPI:         configs.DefaultBotAPI,
		APIKey:         config.APIKey,
		UpdateConfigs:  configs.DefaultUpdateConfigs(),
		Webhook:        false,
		LogFileAddress: configs.DefaultLogFile,
	}

	bot, err := telego.NewBot(telegoCfg)
	if err != nil {
		panic(err)
	}

	if err := bot.Run(); err != nil {
		panic(err)
	}

	return bot
}
