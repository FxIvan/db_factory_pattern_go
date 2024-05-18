package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/fxivan/db_go_abstract_factory/configuration"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client     *mongo.Client
	Database   *mongo.Database
	Collection *mongo.Collection
}

func New(config *configuration.Configuration) (*MongoDB, error) {

	connectionString := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s?authSource=admin", config.User, config.Password, config.Host, config.Port, config.DBName)
	clientOptions := options.Client().ApplyURI(connectionString).SetAuth(options.Credential{
		Username: config.User,
		Password: config.Password,
	})
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("could not connect to MongoDB: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("could not ping MongoDB: %w", err)
	}

	db := client.Database(config.DBName)
	collection := db.Collection("product")

	return &MongoDB{
		Client:     client,
		Database:   db,
		Collection: collection,
	}, nil
}

func (m *MongoDB) Find(id int) string {
	return "function Data | from MongoDB"
}

func (m *MongoDB) Save(data string) error {
	_, err := m.Collection.InsertOne(context.Background(), bson.M{"data": data})
	if err != nil {
		return fmt.Errorf("could not insert into MongoDB: %w", err)
	}
	fmt.Printf("function Save | data to MongoDB %s", data)
	return nil
}
