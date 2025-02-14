package util

import (
	"errors"
	"log"
	"social_network/constant/noti"

	"golang.org/x/crypto/bcrypt"
)

func ToHashString(src string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(src), 10)

	if err != nil {
		log.Println("Error while generating to hash string - " + err.Error())
		return "", errors.New(noti.InternalErr)
	}

	return string(bytes), nil
}

func IsHashStringMatched(inputtedStr, hashedStr string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedStr), []byte(inputtedStr)) == nil
}
