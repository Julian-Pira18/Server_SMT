package main

import (
	"fmt"
	"net/smtp"
	"os"
)

func main() {
	host := "smtp.gmail.com"
	port := "587"
	from := "Your_User@gmail.com"
	password := "Password of Email application"

	toList := []string{"addressee @gmail.com"}
	msg := "Subject: Subject of email \r\n" + "\r\n" + "message email Hi WORLD GMAIL"
	body := []byte(msg)

	auth := smtp.PlainAuth("", from, password, host)

	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Email sent successfully")
}
