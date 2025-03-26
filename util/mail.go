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
	var errLogMsg string = fmt.Sprintf(noti.MailErrMsg, "Util.Mail - SendMail")

	template, err := template.ParseFiles(req.TemplatePath)
	if err != nil {
		log.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}
	req.Body.Subject = req.Subject

	var body bytes.Buffer
	if err := template.Execute(&body, req.Body); err != nil {
		log.Println(errLogMsg + err.Error())
		return errors.New(noti.InternalErr)
	}

	serviceEmail := os.Getenv(env.SERVICE_EMAIL)
	securityPass := os.Getenv(env.SECURITY_PASS)
	host := os.Getenv(env.HOST)
	port, err := strconv.Atoi(os.Getenv(env.MAIL_PORT))
	if err != nil {
		port = 587
	}

	m := mail.NewMessage()
	m.SetHeader("From", serviceEmail)
	m.SetHeader("To", req.Body.Email)
	m.SetHeader("Subject", req.Body.Subject)
	m.SetBody("text/html", body.String())

	dialer := mail.NewDialer(host, port, serviceEmail, securityPass)

	if err := dialer.DialAndSend(m); err != nil {
		log.Println(errLogMsg + err.Error())
		return errors.New(noti.GenerateMailWarnMsg)
	}

	return nil
}
