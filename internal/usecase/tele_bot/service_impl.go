package tele_bot

import (
	"encoding/json"
	"fmt"

	"github.com/SakoDroid/telego"
	"github.com/SakoDroid/telego/objects"

	"github.com/bmkgBot/internal/infrastructure/bmkg"
	"github.com/bmkgBot/internal/shared"
)

type telebotService struct {
	bot         *telego.Bot
	bmkgWrapper bmkg.BMKGWrapperIFace
}

func NewTelebotService(bot *telego.Bot, bmkgWrapper bmkg.BMKGWrapperIFace) TelebotServiceIFace {
	if bot == nil {
		panic("telebot is nil")
	}
	if bmkgWrapper == nil {
		panic("bmkg wrapper is nil")
	}

	return &telebotService{
		bot:         bot,
		bmkgWrapper: bmkgWrapper,
	}
}

func (s *telebotService) EarthquakeInfo() func(u *objects.Update) {
	return func(u *objects.Update) {
		var resp EarthquakeInfoResp

		body, err := s.bmkgWrapper.Get(shared.EarthquakeInfo)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := json.Unmarshal(body, &resp); err != nil {
			fmt.Println(err)
			return
		}

		_, err = s.bot.SendMessage(u.Message.Chat.Id, constructEarthquakeInfoResp(resp.InfoGempa.Gempa), "", u.Message.MessageId, false, false)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (s *telebotService) EarthquakeInfoList() func(u *objects.Update) {
	return func(u *objects.Update) {
		var resp EarthquakeInfoListResp

		body, err := s.bmkgWrapper.Get(shared.EarthquakeInfoList)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err := json.Unmarshal(body, &resp); err != nil {
			fmt.Println(err)
			return
		}

		_, err = s.bot.SendMessage(u.Message.Chat.Id, constructEarthquakeInfoList(resp.InfoGempa.Gempa), "", u.Message.MessageId, false, false)
		if err != nil {
			fmt.Println(err)
		}
	}
}
