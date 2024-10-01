package strategies

import (
	"challenge/internal/models"
	"challenge/internal/repositories"
	"challenge/internal/services/calculation"
	"encoding/json"
	"fmt"
	"time"
)

// JSONProcessor define la estrategia para procesar archivos JSON
type JSONProcessor struct{}

// Data es el formato de los datos en el archivo JSON

// Process procesa el archivo JSON leído desde S3
func (p *JSONProcessor) Process(bucket, key string) error {
	// Leer el archivo desde S3 usando el repositorio
	data, err := repositories.DownloadFileFromS3(bucket, key)
	if err != nil {
		return err
	}

	// Deserializar el JSON
	var jsonData models.FileData
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return fmt.Errorf("error al decodificar JSON: %w", err)
	}
	jsonData.UploadDate = time.Now().Format("02/01/2006 15:04:05")
	jsonData.FileName = key
	emailData, err := calculation.ProcessTransactions(jsonData)
	emailData.File = "JSON"

	if err != nil {
		return fmt.Errorf("error al decodificar JSON: %w", err)
	}

	err = repositories.InsertFileData(jsonData)
	if err != nil {
		return fmt.Errorf("error al guardar datos: %w", err)
	}

	err = repositories.SendNotificationSes(jsonData.Person.Email, emailData)
	if err != nil {
		return err
	}

	err = repositories.SendNotificationSNS(jsonData.Person.Email, emailData)
	if err != nil {
		return err
	}

	return nil
}
