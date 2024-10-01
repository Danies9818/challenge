package models

type Transaction struct {
	ID     string `bson:"id" json:"id"`
	Date   string `bson:"date" json:"date"`
	Amount string `bson:"amount" json:"amount"`
}

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type FileData struct {
	UploadDate   string        `bson:"upload_date" json:"upload_date"`
	Person       Person        `bson:"person" json:"person"`
	FileName     string        `bson:"file_name" json:"file_name"`
	Transactions []Transaction `bson:"transactions" json:"transactions"`
}
