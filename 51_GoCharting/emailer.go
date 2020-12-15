package main

import (
	"log"
	"net/smtp"
	"strings"
)

type email struct {
	socket   string
	password string
	from     string
	mime     string
	subject  string
	body     string
	to       []string
	cc       []string
	message  []byte
}

func sendText(body string) {
	from := "sender@gmail.com"
	pass := "8888"
	to := "recipient@gmail.com"
	msg := "From: " + from + "\n" + "To: " + to + "\n" + "Subject: Hello there\n\n" + body

	errMail := smtp.SendMail("127.0.0.1:1025", smtp.PlainAuth("", from, pass, "127.0.0.1"), from, []string{to}, []byte(msg))
	if errMail != nil {
		log.Printf("smtp error: %s", errMail)
		return
	}
	log.Print("sent email")
}

func sendEmail(emailInfo *email) {
	// to introduce anonymous function to be used in ifs

	if strings.Contains(emailInfo.mime, "Content-Type: text/plain") {
		emailInfo.message = []byte("From: " + emailInfo.from + "\n" + "To: " + strings.Join(emailInfo.to, ",") + "\n" + emailInfo.subject + "\n" + emailInfo.body)
	}
	if strings.Contains(emailInfo.mime, "Content-Type: text/html") {

	}
	if errMail := smtp.SendMail("127.0.0.1:1025", smtp.PlainAuth("", emailInfo.from, emailInfo.password, "127.0.0.1"), emailInfo.from, emailInfo.to, emailInfo.message); errMail != nil {
		log.Printf("smtp error: %s", errMail)
		return
	}
	log.Print("sent email")
}
