package tele_bot

import (
	"reflect"
	"strings"
)

func constructEarthquakeInfoResp(data Earthquake) string {
	var (
		sb  strings.Builder
		val = reflect.ValueOf(data)
	)

	for i := 0; i < val.NumField(); i++ {
		sb.WriteString(val.Type().Field(i).Name + ":\t" + val.Field(i).String() + "\n")
	}

	return sb.String()
}

func constructEarthquakeInfoList(data []Earthquake) string {
	var (
		sb strings.Builder
	)

	for _, v := range data {
		sb.WriteString(constructEarthquakeInfoResp(v) + "\n\n")
	}

	return sb.String()
}
