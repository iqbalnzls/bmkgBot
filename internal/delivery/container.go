package delivery

import (
	telego "github.com/SakoDroid/telego"

	"github.com/bmkgBot/internal/infrastructure/bmkg"
	"github.com/bmkgBot/internal/shared/config"
	restClient "github.com/bmkgBot/internal/shared/rest_client"
	inTelego "github.com/bmkgBot/internal/shared/telego"
	telebot "github.com/bmkgBot/internal/usecase/tele_bot"
)

type Container struct {
	Bot            *telego.Bot
	Config         *config.Config
	TelebotService telebot.TelebotServiceIFace
}

func SetupContainer() *Container {
	cfg := config.NewConfig("./resources/config.yaml")

	client := restClient.NewRestClient(cfg.BMKG.Options)

	bot := inTelego.NewTelego(cfg.Telebot)

	bmkgWrapper := bmkg.NewBMKGWrapperIFace(client, cfg.BMKG)

	telebotService := telebot.NewTelebotService(bot, bmkgWrapper)

	return &Container{
		Bot:            bot,
		Config:         cfg,
		TelebotService: telebotService,
	}
}
