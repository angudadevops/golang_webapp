// forms.go
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
)

type ContactDetails struct {
	FEmail        string
	EmailPassword string
	TEmail        string
	Subject       string
	Message       string
}

func main() {
	tmpl := template.Must(template.ParseFiles("forms.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}

		details := ContactDetails{
			FEmail:        r.FormValue("femail"),
			EmailPassword: r.FormValue("epass"),
			TEmail:        r.FormValue("temail"),
			Subject:       r.FormValue("subject"),
			Message:       r.FormValue("message"),
		}

		// do something with details
		_ = details

		// Address URI to smtp server
		// Sender data.
		from := details.FEmail
		password := details.EmailPassword
		// Receiver email address.
		to := []string{details.TEmail}
		// smtp server configuration.
		smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
		// Message.
		//message := []byte("This is a really unimaginative message, I know.")
		message := []byte("To: " + details.TEmail +
			"\r\n" +
			"Subject: " + details.Subject +
			"\r\n" +
			details.Message)
		//m := "Subject: discount Gophers!\r\n"
		//copy(message[:], details.Message)
		fmt.Printf("%q\n", message)
		// Authentication.
		auth := smtp.PlainAuth(from, from, password, smtpServer.host)
		// Sending email.
		err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Email Sent!")

		tmpl.Execute(w, struct{ Success bool }{true})
	})

	http.ListenAndServe(":8080", nil)
}

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

