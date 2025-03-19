package cmd

import (
	"fmt"
	"log"
	"social_network/constant/noti"

	"github.com/joho/godotenv"
)

func config(logger *log.Logger) {
	if err := godotenv.Load(); err != nil {
		logger.Fatal(fmt.Sprintf(noti.EnvLoadErr, "") + err.Error())
	}
}
