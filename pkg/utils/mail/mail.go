package mail

import (
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

var Env map[string]string
var Auth smtp.Auth

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Env = make(map[string]string)
	Env["SMTP_SERVER"] = os.Getenv("SMTP_SERVER")
	Env["SMTP_USER"] = os.Getenv("SMTP_USER")
	Env["SMTP_PASS"] = os.Getenv("SMTP_PASS")
	Env["SMTP_SSL_PORT"] = os.Getenv("SMTP_SSL_PORT")
	Env["SMTP_TLS_PORT"] = os.Getenv("SMTP_TLS_PORT")

	if Env["SMTP_USER"] == "" || Env["SMTP_PASS"] == "" || Env["SMTP_SERVER"] == "" {
		log.Fatal("SMTP environment variables are not properly set")
	}
}

// Configure SMTP authentication
func Builder() {
	Auth = smtp.PlainAuth("", Env["SMTP_USER"], Env["SMTP_PASS"], Env["SMTP_SERVER"])
}

// Send email
func SendEmail(msg []byte, to []string) {
	sender := Env["SMTP_USER"]
	smtpPort := Env["SMTP_TLS_PORT"]

	err := smtp.SendMail(Env["SMTP_SERVER"]+":"+smtpPort, Auth, sender, to, msg)
	if err != nil {
		log.Fatalf("Failed to send email: %v", err)
	} else {
		log.Println("Email sent successfully!")
	}
}
