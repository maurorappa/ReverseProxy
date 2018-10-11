package main

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"net/mail"
	"net/smtp"
)

var (
	letters = []rune("abcdefghijklmnopqrstuvwxyz_-!.=$@:1234567890")
)

func inArray(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func send_email(conf *config,code string) {

	from := mail.Address{"", conf.Mailbox}
	to := mail.Address{"", conf.Mailbox}
	subj := "Reverse Proxy Urlknocking codes"
	body := "This is an important code to protect your blog.\n" + code + "\n\n"

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := conf.Smtp_server

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", conf.Auth_user, conf.Auth_pwd, host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         host,
	}

	c, err := smtp.Dial(servername)
	if err != nil {
		fmt.Printf("%s /n", err)
	}

	c.StartTLS(tlsconfig)

	// Auth
	if err = c.Auth(auth); err != nil {
		fmt.Printf("%s /n", err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		fmt.Printf("%s /n", err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		fmt.Printf("%s /n", err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		fmt.Printf("%s /n", err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		fmt.Printf("%s /n", err)
	}

	err = w.Close()
	if err != nil {
		fmt.Printf("%s /n", err)
	}

	c.Quit()

}
