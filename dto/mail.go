package dto

import "log"

type MailBody struct {
	Email    string
	Password string
	Username string
	Url      string
}

type SendMailRequest struct {
	Body         MailBody    `json:"mail_body"`
	TemplatePath string      `json:"template_path"`
	Subject      string      `json:"subject"`
	Logger       *log.Logger `json:"logger"`
}
