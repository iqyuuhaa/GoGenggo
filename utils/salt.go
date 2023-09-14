package utils

import (
	"fmt"
	"math/rand"

	"gogenggo/internals/types/constants"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateRandomString() string {
	s := make([]byte, 4)
	for i := range s {
		s[i] = constants.LetterRunes[rand.Intn(len(constants.LetterRunes))]
	}

	return string(s)
}

func AddSaltingSuffix(s string) (str string) {
	genString := GenerateRandomString()
	str = fmt.Sprintf("%s%s", s, genString)
	return
}

func RemoveSaltingSuffix(s string) (str string) {
	str = s[:len(s)-4]
	return
}

func AddSaltingPrefix(s string) (str string) {
	genString := GenerateRandomString()
	str = fmt.Sprintf("%s%s", genString, s)
	return
}

func RemoveSaltingPrefix(s string) (str string) {
	str = s[4:]
	return
}
