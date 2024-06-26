package taskhandler

import (
	"errors"
	"net/smtp"
	"os"
)

func SendMailTask(receiver string) error {
	// sender data
	from, ok := os.LookupEnv("SENDER_MAIL")
	if !ok {
		return errors.New("env SENDER_MAIL not set")
	}
	password, ok := os.LookupEnv("SENDER_PASSWORD")
	if !ok {
		return errors.New("env SENDER_PASSWORD not set")
	}
	senderUserName, ok := os.LookupEnv("SENDER_USERNAME")
	if !ok {
		return errors.New("env SENDER_USERNAME not set")
	}
	// receiver address
	toEmail := receiver // ex: "Jane.Smith@yahoo.com"
	to := []string{toEmail}
	// smtp - Simple Mail Transfer Protocol
	host, ok := os.LookupEnv("SENDER_HOST")
	if !ok {
		return errors.New("env SENDER_HOST not set")
	}
	port, ok := os.LookupEnv("SMTP_PORT")
	if !ok {
		return errors.New("env SMTP_PORT not set")
	}
	address := host + ":" + port
	// Set up authentication information.
	auth := smtp.PlainAuth("", senderUserName, password, host)
	msg := []byte(
		"From: " + from + " <" + from + ">\r\n" +
			"To: " + toEmail + "\r\n" +
			"Subject: Greetings Message !\r\n" +
			"MIME: MIME-version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n" + "\r\n" + `Hi, How Are You !`)
	return smtp.SendMail(address, auth, from, to, msg)
}
