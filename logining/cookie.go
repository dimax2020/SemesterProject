package logining

import "time"

func CheckCookie() {

}

func GenerateNewCookie() (string, time.Time) {
	return "Cookie", time.Now().Add(8640 * time.Hour)
}