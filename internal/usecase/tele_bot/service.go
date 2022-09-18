package tele_bot

import "github.com/SakoDroid/telego/objects"

type TelebotServiceIFace interface {
	EarthquakeInfo() func(u *objects.Update)
	EarthquakeInfoList() func(u *objects.Update)
}
