package util

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"social_network/constant/env"
	"social_network/constant/noti"
	"social_network/dto"
	"strconv"
	"text/template"

	"gopkg.in/mail.v2"
)

func SendMail(req dto.SendMailRequest) error {
	var body bytes.Buffer

	template, err := template.ParseFiles(req.TemplatePath)
	var errLogMsg string = fmt.Sprintf(noti.MailErrMsg, "SendMail")

	if err != nil {
		log.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}
	template.Execute(&body, req.Body)

	var serviceEmail string = os.Getenv(env.SERVICE_EMAIL)
	var securityPass string = os.Getenv(env.SECURITY_PASS)
	var host string = os.Getenv(env.HOST)
	port, err := strconv.Atoi(os.Getenv(env.MAIL_PORT))
	if err != nil {
		port = 587
	}

	var m = mail.NewMessage()
	m.SetHeader("From", serviceEmail)
	m.SetHeader("To", req.Body.Email)
	m.SetHeader("Subject", req.Subject)
	m.SetBody("text/html", body.String())

	diabler := mail.NewDialer(host, port, serviceEmail, securityPass)

	if err := diabler.DialAndSend(m); err != nil {
		log.Println(errLogMsg + err.Error())
		return errors.New(noti.GenerateMailWarnMsg)
	}

	return nil
}
