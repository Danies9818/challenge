package strategies

import (
	"bytes"
	"challenge/internal/models"
	"challenge/internal/repositories"
	"challenge/internal/services/calculation"
	"encoding/csv"
	"fmt"
	"time"
)

type CSVProcessor struct{}

func (p *CSVProcessor) Process(bucket, key string) error {
	data, err := repositories.DownloadFileFromS3(bucket, key)
	if err != nil {
		return err
	}

	var transactions []models.Transaction
	var person models.Person

	reader := bytes.NewReader(data)

	csvReader := csv.NewReader(reader)

	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Error leyendo el CSV:", err)
		return err
	}

	for i, record := range records {
		if i == 0 {
			continue
		}

		transaction := models.Transaction{
			ID:     record[0],
			Date:   record[1],
			Amount: record[2],
		}

		person = models.Person{
			Name:  record[3],
			Email: record[4],
		}
		// Agregar la transacción a la lista
		transactions = append(transactions, transaction)

	}

	dataFile := models.FileData{
		Person:       person,
		Transactions: transactions,
		FileName:     key,
		UploadDate:   time.Now().Format("02/01/2006 15:04:05"),
	}

	emailData, err := calculation.ProcessTransactions(dataFile)
	emailData.File = "CSV"

	if err != nil {
		return fmt.Errorf("error al decodificar CSV: %w", err)
	}

	err = repositories.InsertFileData(dataFile)
	if err != nil {
		return fmt.Errorf("error al guardar datos: %w", err)
	}

	// Procesar los datos
	err = repositories.SendNotificationSes(dataFile.Person.Email, emailData)
	if err != nil {
		return err
	}

	err = repositories.SendNotificationSNS(dataFile.Person.Email, emailData)
	if err != nil {
		return err
	}

	return nil
}
