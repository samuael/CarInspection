package main

import (
	"log"
	"net/smtp"
)

func main() {
	// Choose auth method and set it up
	auth := smtp.PlainAuth("", "samuaeladnew@gmail.com", "0774samuael", "smtp.gmail.com:465")

	// Here we do it all: connect to our server, set up a message and send it
	to := []string{"samuaeladnew.zebir@gmail.com"}
	msg := []byte("To: bill@gates.com\r\n" +
		"Subject: Why are you not using Mailtrap yet? Here’s the space for our great sales pitch")
	err := smtp.SendMail("smtp.gmail.com:465", auth, "samuaeladnew@gmail.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
}
