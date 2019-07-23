package qasircore

import (
	"fmt"

	"github.com/mostafah/mandrill"
)

type Message struct {
	ToEmail      string
	ToName       string
	FromEmail    string
	FromName     string
	Subject      string
	PathTemplate string
	TemplateData map[string]string
	Text         string
}

type mail struct {
	msg         Message
	configEmail map[string]interface{}
}

func (this *mail) SetMessage(msg Message) {
	this.msg = msg
}

func (this *mail) SetConfig(config map[string]interface{}) {
	this.configEmail = config
}

func (this *mail) Send(msg Message) error {
	driverEmail := fmt.Sprint(this.configEmail["driver"])
	this.msg = msg
	if driverEmail == "mandrill" {
		mandrillMessage := mandrill.NewMessageTo(this.msg.ToEmail, this.msg.ToName)
		mandrillMessage.FromEmail = this.msg.FromEmail
		mandrillMessage.FromName = this.msg.FromName
		mandrillMessage.Subject = this.msg.Subject

		_, err := mandrillMessage.SendTemplate(this.msg.PathTemplate, this.msg.TemplateData, false)

		if err != nil {
			return err
		}
	}
	return nil
}

func Mail(config map[string]interface{}) *mail {
	var email mail
	email.SetConfig(config)
	return &email
}
