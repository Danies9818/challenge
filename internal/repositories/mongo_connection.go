package repositories

import (
	"context"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	clientInstance    *mongo.Client
	clientInstanceErr error
	mongoOnce         sync.Once
	mongoURI          = os.Getenv("MONGODB_URI")
	mongoDatabase     = os.Getenv("MONGODB_DATABASE")
	mongoCollection   = os.Getenv("MONGODB_COLLECTION")
)

// GetMongoClient retorna una instancia singleton del cliente de MongoDB
func GetMongoClient() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		clientOptions := options.Client().ApplyURI(mongoURI)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			clientInstanceErr = err
			return
		}

		// Verificar la conexión
		err = client.Ping(ctx, nil)
		if err != nil {
			clientInstanceErr = err
			return
		}

		clientInstance = client
	})

	return clientInstance, clientInstanceErr
}

// InsertFileData inserta un documento en la colección de MongoDB
func InsertFileData(fileData interface{}) error {
	client, err := GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(mongoDatabase).Collection(mongoCollection)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, fileData)
	if err != nil {
		return err
	}

	return nil
}
