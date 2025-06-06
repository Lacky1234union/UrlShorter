package random

import "github.com/labstack/gommon/random"

func NewRandomString(length int) string {
	return random.String(uint8(length), "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890")
}
