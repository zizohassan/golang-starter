package helpers

import (
	"fmt"
	"net/smtp"
	"os"
)

/**
* function to send email with subject and content
 */
func SendMail(email string, subject string , content string) {
	go func() {
		auth := smtp.PlainAuth("", os.Getenv("STMP_EMAIL_ADDRESS"), os.Getenv("STMP_EMAIL_PASSWORD"), os.Getenv("STMP_EMAIL_HOST"))
		to := []string{email}
		msg := []byte("To : " + email + "\r\n" + "Subject : " +subject + "\r\n" + "\r\n" + content + "\r\n" )
		fmt.Println(os.Getenv("STMP_EMAIL_ADDRESS")+":"+os.Getenv("STMP_EMAIL_PORT"))
		err := smtp.SendMail(os.Getenv("STMP_EMAIL_ADDRESS")+":"+os.Getenv("STMP_EMAIL_PORT"), auth, os.Getenv("STMP_EMAIL_ADDRESS"), to, msg)
		fmt.Println("++++++++++++")
		fmt.Println(err.Error())
		fmt.Println("++++++++++++")
	}()
	return
}
