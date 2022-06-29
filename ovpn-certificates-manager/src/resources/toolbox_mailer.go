package resources

import (
	"log"
	"strings"
)

type mailer struct {

}

func (m mailer) SendEmail(to, attachmentName, attachmentContent string) (succeed bool) {

	if Session.Switcher.sendEmail {

		name := strings.Title(strings.Split(to, ".")[0])

		err, message := Session.Config.Mailer.EmailMessage.ParseTemplate(
			struct {
				Name string
			}{
				Name: name,
			})

		if err != nil {
			log.Printf("ERROR: Error parsing email template for %s:\n%v", to, err)
			return false
		}

		mailAgent := Session.Config.Mailer.Smtp.GetMailAgent()

		mailAgent.To(to)
		mailAgent.From(Session.Config.Mailer.Smtp.Username)
		mailAgent.Subject(Session.Config.Mailer.EmailMessage.Subject)
		mailAgent.HTML().Set(message)

		mailAgent.Attach(attachmentName, strings.NewReader(attachmentContent))

		if err := mailAgent.Send(); err != nil {
			log.Printf("ERROR: Error sending email to %s:\n%v", to, err)
			return false
		}

		return true
	} else {
		return false
	}
}

