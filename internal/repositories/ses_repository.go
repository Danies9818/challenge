package repositories

import (
	"bytes"
	"challenge/internal/services/calculation"
	"fmt"
	"os"
	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

const (
	sender  = "danies9818@gmail.com"
	subject = "Notificación de Transacciones Procesadas"
	charset = "UTF-8"
)

func SendNotificationSes(email string, data calculation.EmailData) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))
	svc := ses.New(sess)

	templateFile := "templates/email_template.html"
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("error al cargar la plantilla de correo: %v", err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("error al procesar la plantilla de correo: %v", err)
	}

	input := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(email),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Charset: aws.String(charset),
					Data:    aws.String(body.String()),
				},
			},
			Subject: &ses.Content{
				Charset: aws.String(charset),
				Data:    aws.String(subject),
			},
		},
		Source: aws.String(sender),
	}

	_, err = svc.SendEmail(input)
	if err != nil {
		return fmt.Errorf("error al enviar el correo electrónico: %v", err)
	}

	fmt.Println("Correo enviado a través de SES a:", email)
	return nil
}
