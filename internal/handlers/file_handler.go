package handlers

import (
	"challenge/internal/services"
	"context"
	"fmt"
	"path/filepath"

	"github.com/aws/aws-lambda-go/events"
)

func HandleS3Event(ctx context.Context, s3Event events.S3Event) error {
	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		fmt.Printf("Procesando archivo subido: Bucket: %s, Key: %s\n", bucket, key)
		ext := filepath.Ext(key)
		strategyContext := services.NewStrategyContext(ext)
		err := strategyContext.Execute(bucket, key)
		if err != nil {
			fmt.Printf("Error al procesar el archivo: %s\n", err)
			return err
		}
	}
	return nil
}
