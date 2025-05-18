package mail

import (
	"crypto/tls"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
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
	Auth = smtp.PlainAuth("", Env["SMTP_MAIL"], Env["SMTP_PASS"], Env["SMTP_SERVER"])
}

func SendEmailTLS(msg []byte, to []string) {
	smtpServer := Env["SMTP_SERVER"] + ":" + Env["SMTP_TLS_PORT"]

	// Connect to SMTP server
	conn, err := smtp.Dial(smtpServer)
	if err != nil {
		log.Fatalf("Error connecting to SMTP server: %v", err)
	}
	defer conn.Close()

	// EHLO greeting
	if err = conn.Hello(Env["SMTP_SERVER"]); err != nil {
		log.Fatalf("Error during EHLO: %v", err)
	}

	// Start TLS
	if ok, _ := conn.Extension("STARTTLS"); ok {
		if err = conn.StartTLS(&tls.Config{ServerName: Env["SMTP_SERVER"]}); err != nil {
			log.Fatalf("Error during STARTTLS: %v", err)
		}
	} else {
		log.Fatal("Server does not support STARTTLS")
	}

	// SMTP authentication
	Auth := smtp.PlainAuth("", Env["SMTP_MAIL"], Env["SMTP_PASS"], Env["SMTP_SERVER"])
	if err = conn.Auth(Auth); err != nil {
		log.Fatalf("SMTP authentication error: %v", err)
	}

	// Set sender
	if err = conn.Mail(Env["SMTP_MAIL"]); err != nil {
		log.Fatalf("Error setting sender: %v", err)
	}

	// Check that there is at least one recipient
	if len(to) == 0 {
		log.Fatal("No recipients to send the email")
	}

	// Set recipients
	for _, recipient := range to {
		if err = conn.Rcpt(recipient); err != nil {
			log.Fatalf("Error with recipient %s: %v", recipient, err)
		}
	}

	// Get writer to send data
	w, err := conn.Data()
	if err != nil {
		log.Fatalf("Error starting data send: %v", err)
	}

	// Write message
	_, err = w.Write(msg)
	if err != nil {
		log.Fatalf("Error writing message: %v", err)
	}

	// Close writer to finalize DATA
	err = w.Close()
	if err != nil {
		log.Fatalf("Error closing data writer: %v", err)
	}

	// Properly close SMTP connection
	err = conn.Quit()
	if err != nil {
		log.Fatalf("Error closing SMTP connection: %v", err)
	}

	log.Println("Email sent successfully using STARTTLS on port 587!")
}

// Send email using SSL on port 465
func SendEmailSSL(msg []byte, to []string) {
	smtpServer := Env["SMTP_SERVER"] + ":" + Env["SMTP_SSL_PORT"]

	// Configure TLS with SMTP server
	tlsConfig := &tls.Config{
		ServerName: Env["SMTP_SERVER"],
	}

	// Establish secure TLS connection
	conn, err := tls.Dial("tcp", smtpServer, tlsConfig)
	if err != nil {
		log.Fatalf("Error establishing SSL connection: %v", err)
	}
	defer conn.Close()

	// Create SMTP client over TLS connection
	client, err := smtp.NewClient(conn, Env["SMTP_SERVER"])
	if err != nil {
		log.Fatalf("Error creating SMTP client: %v", err)
	}
	defer client.Quit() // Close SMTP session at the end

	// SMTP authentication
	auth := smtp.PlainAuth("", Env["SMTP_MAIL"], Env["SMTP_PASS"], Env["SMTP_SERVER"])
	if err = client.Auth(auth); err != nil {
		log.Fatalf("SMTP authentication error: %v", err)
	}

	// Set sender
	if err = client.Mail(Env["SMTP_MAIL"]); err != nil {
		log.Fatalf("Error setting sender: %v", err)
	}

	// Validate recipients exist
	if len(to) == 0 {
		log.Fatal("No recipients to send the email")
	}

	// Set recipients
	for _, recipient := range to {
		if err = client.Rcpt(recipient); err != nil {
			log.Fatalf("Error with recipient %s: %v", recipient, err)
		}
	}

	// Get writer to send message body
	w, err := client.Data()
	if err != nil {
		log.Fatalf("Error starting data send: %v", err)
	}

	// Write message
	_, err = w.Write(msg)
	if err != nil {
		log.Fatalf("Error writing message: %v", err)
	}

	// Close writer to finalize DATA
	err = w.Close()
	if err != nil {
		log.Fatalf("Error closing data writer: %v", err)
	}

	// Properly close SMTP session
	if err = client.Quit(); err != nil {
		log.Fatalf("Error closing SMTP connection: %v", err)
	}

	log.Println("Email sent successfully using SSL on port 465!")
}
