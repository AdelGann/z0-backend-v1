package mail

import (
	"crypto/tls"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/smtp"
	"os"
)

var Env map[string]string
var Auth smtp.Auth

// Load environment variables
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Env = make(map[string]string)
	Env["SMTP_SERVER"] = os.Getenv("SMTP_SERVER")
	Env["SMTP_MAIL"] = os.Getenv("SMTP_MAIL")
	Env["SMTP_USER"] = os.Getenv("SMTP_USER")
	Env["SMTP_PASS"] = os.Getenv("SMTP_PASS")
	Env["SMTP_SSL_PORT"] = os.Getenv("SMTP_SSL_PORT")
	Env["SMTP_TLS_PORT"] = os.Getenv("SMTP_TLS_PORT")

	if Env["SMTP_MAIL"] == "" || Env["SMTP_PASS"] == "" || Env["SMTP_SERVER"] == "" {
		log.Fatal("SMTP environment variables are not properly configured")
	}
}

// Configure SMTP authentication
func Builder() {
	LoadEnv()
	Auth = smtp.PlainAuth("", Env["SMTP_MAIL"], Env["SMTP_PASS"], Env["SMTP_SERVER"])
}

func SendEmailTLS(msg []byte, to []string) error {
	smtpServer := Env["SMTP_SERVER"] + ":" + Env["SMTP_TLS_PORT"]

	conn, err := smtp.Dial(smtpServer)
	if err != nil {
		return err
	}
	defer conn.Close()

	if err = conn.Hello(Env["SMTP_SERVER"]); err != nil {
		return err
	}

	if ok, _ := conn.Extension("STARTTLS"); ok {
		if err = conn.StartTLS(&tls.Config{ServerName: Env["SMTP_SERVER"]}); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("server does not support STARTTLS")
	}

	auth := smtp.PlainAuth("", Env["SMTP_MAIL"], Env["SMTP_PASS"], Env["SMTP_SERVER"])
	if err = conn.Auth(auth); err != nil {
		return err
	}

	if err = conn.Mail(Env["SMTP_MAIL"]); err != nil {
		return err
	}

	if len(to) == 0 {
		return fmt.Errorf("no recipients to send the email")
	}

	for _, recipient := range to {
		if err = conn.Rcpt(recipient); err != nil {
			return fmt.Errorf("error with recipient %s: %v", recipient, err)
		}
	}

	w, err := conn.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	err = conn.Quit()
	if err != nil {
		return err
	}

	log.Println("Email sent successfully using STARTTLS on port 587!")
	return nil
}

func SendEmailSSL(msg []byte, to []string) error {
	smtpServer := Env["SMTP_SERVER"] + ":" + Env["SMTP_SSL_PORT"]

	tlsConfig := &tls.Config{
		ServerName: Env["SMTP_SERVER"],
	}

	conn, err := tls.Dial("tcp", smtpServer, tlsConfig)
	if err != nil {
		return err
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, Env["SMTP_SERVER"])
	if err != nil {
		return err
	}
	defer client.Quit()

	auth := smtp.PlainAuth("", Env["SMTP_MAIL"], Env["SMTP_PASS"], Env["SMTP_SERVER"])
	if err = client.Auth(auth); err != nil {
		return err
	}

	if err = client.Mail(Env["SMTP_MAIL"]); err != nil {
		return err
	}

	if len(to) == 0 {
		return fmt.Errorf("no recipients to send the email")
	}

	for _, recipient := range to {
		if err = client.Rcpt(recipient); err != nil {
			return fmt.Errorf("error with recipient %s: %v", recipient, err)
		}
	}

	w, err := client.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	if err = client.Quit(); err != nil {
		return err
	}

	log.Println("Email sent successfully using SSL on port 465!")
	return nil
}
