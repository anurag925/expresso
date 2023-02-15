package initializers

import (
	"context"

	"github.com/anurag925/mongoboiler"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo interface {
	DB() *mongoboiler.DB
	Connect(dbName string) (*mongoboiler.DB, error)
	Disconnect() error
}

// db := os.Getenv("MONGO_DB_NAME")
type MongoDB struct {
	ctx context.Context
	url string
	db  *mongoboiler.DB
}

var _ Mongo = (*MongoDB)(nil)

func NewMongoDB(ctx context.Context, url string, db *mongoboiler.DB) *MongoDB {
	return &MongoDB{ctx, url, db}
}

func (m *MongoDB) DB() *mongoboiler.DB {
	return m.db
}

func (m *MongoDB) Connect(dbName string) (*mongoboiler.DB, error) {
	client, err := mongo.Connect(m.ctx, options.Client().ApplyURI(m.url))
	if err != nil {
		return nil, err
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	m.db = mongoboiler.New(client, dbName, m.ctx)
	return m.db, nil
}

func (m *MongoDB) Disconnect() error {
	return m.db.Disconnect()
}
