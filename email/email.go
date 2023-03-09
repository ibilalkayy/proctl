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

func Verify(values [5]string) {
	mail := gomail.NewMessage()
	myEmail := middleware.LoadEnvVariable("APP_EMAIL")
	myPassword := middleware.LoadEnvVariable("APP_PASSWORD")

	body := new(bytes.Buffer)
	temp, err := template.ParseFiles("cmd/views/" + values[0] + ".html")
	if err != nil {
		log.Fatal(err)
	}

	getAccountName := AccountInfo{
		GetAccountName: values[1],
		GetEncodedText: values[2],
	}

	if err := temp.Execute(body, getAccountName); err != nil {
		fmt.Println(errors.New("Cannot load the template"))
	}

	mail.SetHeader("From", myEmail)
	mail.SetHeader("To", values[3])
	mail.SetHeader("Reply-To", myEmail)
	mail.SetHeader("Subject", values[4])
	mail.SetBody("text/html", body.String())
	a := gomail.NewDialer("smtp.gmail.com", 587, myEmail, myPassword)
	if err := a.DialAndSend(mail); err != nil {
		log.Fatal(err)
	}
}
