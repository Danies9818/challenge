package repositories

import (
	"bytes"
	"challenge/internal/services/calculation"
	"fmt"
	"os"
	"text/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func SendNotificationSNS(email string, data calculation.EmailData) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	}))
	svc := sns.New(sess)

	templateFile := "templates/email_template.txt"
	tmpl, err := template.ParseFiles(templateFile)
	if err != nil {
		return fmt.Errorf("error al cargar la plantilla de correo: %v", err)
	}

	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return fmt.Errorf("error al procesar la plantilla de correo: %v", err)
	}

	_, err = svc.Publish(&sns.PublishInput{
		Message:  aws.String(body.String()),
		Subject:  aws.String("Transacción procesada"),
		TopicArn: aws.String(os.Getenv("SNS_TOPIC_ARN")),
		MessageAttributes: map[string]*sns.MessageAttributeValue{
			"email": {
				DataType:    aws.String("String"),
				StringValue: aws.String(email),
			},
		},
	})

	fmt.Println("Correo enviado a través de SNS a:", email)
	if err != nil {
		return fmt.Errorf("error al enviar el mensaje SNS: %v", err)
	}

	return nil
}
