package main

import "time"

var emailSent = make(map[string]emailSentValue)

type emailSentValue struct {
	AuthCode string
	Expiry   time.Time
}

func SendEmail(email string) {
	authcode := RandomCode()
	expiry := time.Now().Add(10 * time.Minute)
	emailSent[email] = emailSentValue{
		AuthCode: authcode,
		Expiry:   expiry,
	}
}
