package delivery

import (
	"fmt"

	"github.com/SakoDroid/telego"
	"github.com/SakoDroid/telego/objects"

	telebot "github.com/bmkgBot/internal/usecase/tele_bot"
)

type telebotHandler struct {
	bot            *telego.Bot
	telebotService telebot.TelebotServiceIFace
}

func NewTelebotHandler(bot *telego.Bot, telebotService telebot.TelebotServiceIFace) *telebotHandler {
	if bot == nil {
		panic("telebot is nil")
	}
	if telebotService == nil {
		panic("telebot service is nil")
	}

	return &telebotHandler{
		bot:            bot,
		telebotService: telebotService,
	}
}

func (h *telebotHandler) TelebotHandler() {
	updateChannel := *h.bot.GetUpdateChannel()

	h.bot.AddHandler("info-gempa", h.telebotService.EarthquakeInfo(), "private")
	h.bot.AddHandler("daftar-info-gempa", h.telebotService.EarthquakeInfoList(), "private")

	update := <-updateChannel
	h.validateCommand(update)
	fmt.Println(update)
}

func (h *telebotHandler) validateCommand(u *objects.Update) {
	var m = map[string]string{
		"info-gempa":        "info-gempa",
		"daftar-info-gempa": "daftar-info-gempa",
	}

	if _, ok := m[u.Message.Text]; !ok {
		_, err := h.bot.SendMessage(u.Message.Chat.Id, "Perintah yang anda minta tidak ditemukan, perintah yang dapat dilakukan:\ninfo-gempa: untuk menampilkan info gempa terbaru\ndaftar-info-gempa: untuk menampilkan daftar 15 info gempa terbaru", "", u.Message.MessageId, false, false)
		if err != nil {
			fmt.Println(err)
		}
	}
}
