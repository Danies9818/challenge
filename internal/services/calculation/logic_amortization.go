package calculation

import (
	"challenge/internal/models"
	"fmt"
	"strconv"
	"time"
)

type EmailData struct {
	Person              models.Person
	Transactions        []models.Transaction
	TotalBalance        float64
	TransactionsByMonth map[string]int
	AverageDebitAmount  float64
	AverageCreditAmount float64
	File                string
}

func ProcessTransactions(data models.FileData) (EmailData, error) {

	balance := calculateTotalBalance(data.Transactions)
	transactionsByMonth := countTransactionsByMonth(data.Transactions)
	avgDebit, avgCredit := calculateAverage(data.Transactions)

	results := EmailData{
		Person:              data.Person,
		Transactions:        data.Transactions[:5],
		TotalBalance:        balance,
		TransactionsByMonth: transactionsByMonth,
		AverageDebitAmount:  avgDebit,
		AverageCreditAmount: avgCredit,
	}

	return results, nil
}

func calculateTotalBalance(transactions []models.Transaction) float64 {
	var totalBalance float64
	for _, t := range transactions {
		totalBalance += castAmount(t.Amount)
	}
	return totalBalance
}

func countTransactionsByMonth(transactions []models.Transaction) map[string]int {
	transactionsByMonth := make(map[string]int)
	for _, t := range transactions {
		// Convertir la fecha de la transacción a time.Time
		transactionDate, _ := time.Parse("02/01/2006", t.Date)

		// Formatear el mes como "Enero 2023", "Febrero 2023", etc.
		monthYear := transactionDate.Format("January 2006")

		// Incrementar el contador para ese mes
		transactionsByMonth[monthYear]++
	}
	return transactionsByMonth
}

func calculateAverage(transactions []models.Transaction) (float64, float64) {
	var sumDebit, sumCredit float64
	var countDebit, countCredit int = 0, 0
	var avgDebit, avgCredit = 0.0, 0.0

	for _, t := range transactions {

		num := castAmount(t.Amount)

		if num > 0 {
			sumCredit += num
			countCredit++
		} else {
			sumDebit += num
			countDebit++
		}
	}
	if countDebit == 0 {
		avgDebit = 0
	}
	if avgCredit == 0 {
		avgDebit = 0
	}

	avgDebit = sumDebit / float64(countDebit)
	avgCredit = sumCredit / float64(sumCredit)

	return avgDebit, avgCredit
}

func castAmount(num string) float64 {
	numCast, err := strconv.ParseFloat(num, 64)
	if err != nil {
		fmt.Println("Error al convertir:", err)
		return 0.0
	}
	return numCast
}
