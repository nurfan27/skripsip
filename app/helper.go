package app

import "qasircore"

var (
	Env = qasircore.Env("./")
)

func RemoveCountryCode(phoneTemp string) string {
	phoneTemp = phoneTemp[2:len(phoneTemp)]
	phoneTemp = "0" + phoneTemp

	return phoneTemp
}
