package utils

import (
	"fmt"
	"time"
)

var travel = time.Date(2025, time.January, 1, 00, 0, 0, 0, time.UTC)
var url = "https://api.twitter.com/1.1/account/update_profile.json?name="
var canadianFlag = "%F0%9F%87%A8%F0%9F%87%A6"
var pseudo = "Bivouac%20-%20"

func DaysBeforeCanada() int {
	year, month, day := time.Now().Date()

	today := time.Date(year, month, day, 00, 0, 0, 0, time.UTC)

	return int(travel.Sub(today).Hours() / 24)
}

func GeneratePseudo(days int) string {

	return fmt.Sprintf("%s%dJ%%20%s", pseudo, days, canadianFlag)
}

func GeneratePath(pseudo string) string {
	return url + pseudo
}
