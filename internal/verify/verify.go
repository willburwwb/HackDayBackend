package verify

import (
	"math/rand"
	"time"
)

func init() {
	rand.NewSource(time.Hour.Microseconds())
}

var (
	letters = []rune("0123456")
)

func CreateVerifyCode() string {
	var letters = []byte("123456789")
	verifyCode := make([]byte, 6) // 六位
	for i := range verifyCode {
		verifyCode[i] = letters[rand.Intn(len(letters))] //随机数
	}
	return string(verifyCode)
}

func CreateRandomName() string {
	temp := make([]rune, 6)
	for i := range temp {
		temp[i] = letters[rand.Intn(len(letters))]
	}
	return "user" + string(temp)
}
