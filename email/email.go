package email

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"

	"github.com/ibilalkayy/proctl/middleware"
	"gopkg.in/gomail.v2"
)

type AccountInfo struct {
	GetAccountName string
	GetEncodedText string
}

// Verify sends a verification email
func Verify(values [5]string) {
	mail := gomail.NewMessage()

	// Load email credentials from environment variables
	myEmail := middleware.LoadEnvVariable("APP_EMAIL")
	myPassword := middleware.LoadEnvVariable("APP_PASSWORD")

	body := new(bytes.Buffer)
	temp, err := template.ParseFiles("cmd/views/" + values[0] + ".html")
	if err != nil {
		log.Fatal(err)
	}

	// Create an AccountInfo struct to pass to the email template
	getAccountName := AccountInfo{
		GetAccountName: values[1],
		GetEncodedText: values[2],
	}

	if err := temp.Execute(body, getAccountName); err != nil {
		fmt.Println(errors.New("Cannot load the template"))
	}

	// Set email headers and content
	mail.SetHeader("From", myEmail)
	mail.SetHeader("To", values[3])
	mail.SetHeader("Reply-To", myEmail)
	mail.SetHeader("Subject", values[4])
	mail.SetBody("text/html", body.String())

	// Create a new SMTP dialer using the email credentials
	// and send the email
	dialer := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)
	if err := dialer.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}
}
