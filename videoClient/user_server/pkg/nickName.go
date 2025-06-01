package pkg

import (
	"math/rand"
	"time"
)

const charset = "用户" + "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"
const Fleeting = "瞬影" + "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	a := make([]byte, length)
	for i := range a {
		a[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(a)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
}
func RandomFleetingString(length int) string {
	return StringWithCharset(length, Fleeting)
}
