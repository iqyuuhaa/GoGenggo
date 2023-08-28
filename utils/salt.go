package utils

import (
	"fmt"
	"math/rand"

	"gogenggo/config"
	"gogenggo/internals/types/constants"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	return uuid.New().String()
}

func GenerateRandomString() string {
	s := make([]byte, config.Configs.Main.Generator.SaltLimit)
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
	str = s[:len(s)-config.Configs.Main.Generator.SaltLimit]
	return
}

func AddSaltingPrefix(s string) (str string) {
	genString := GenerateRandomString()
	str = fmt.Sprintf("%s%s", genString, s)
	return
}

func RemoveSaltingPrefix(s string) (str string) {
	str = s[config.Configs.Main.Generator.SaltLimit:]
	return
}
