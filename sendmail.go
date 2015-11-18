package myconfig

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
)

/*
Sendmail
	Takes c for Configuration
	lowerSecurity as bool to deactivate warnings at self-signed
	Set headers:
		From = name
		To = name
		Subject = c.Subject
		Body = text/plain, c.Body
		Attach c.Filename

		return err on error

*/
func Sendmail(c Configuration, lowerSecurity bool) error {
	for _, name := range c.Recipients {

		msg := gomail.NewMessage()
		msg.SetHeader("From", name)
		msg.SetHeader("To", name)
		msg.SetHeader("Subject", c.Subject)
		msg.SetBody("text/plain", c.Body)
		msg.Attach(c.Filename)

		d := gomail.NewPlainDialer(c.MaiHost, c.MailPort, c.MailUsername, c.MailPassword)

		if lowerSecurity {
			d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
		}
		if err := d.DialAndSend(msg); err != nil {
			return err
		}

	}
	return nil
}
