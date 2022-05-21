package mongo

import (
	"context"
	"errors"
	"github.com/crawlab-team/go-trace"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var (
	defaultDbName = "crawlab_db"
)

type mongoDbDatabase struct {
	dbName string
	client *mongo.Client // may be connection pool
	db     *mongo.Database
	cols   sync.Map
}

func NewMongoDbDatabase(name string, client *mongo.Client) *mongoDbDatabase {
	if name == "" {
		name = defaultDbName
	}
	db := &mongoDbDatabase{dbName: name, client: client}
	db.db = client.Database(name)
	return db
}

func (db *mongoDbDatabase) GetClient() *mongo.Client {
	return db.client
}

func (db *mongoDbDatabase) GetColByName(colName string) *Col {
	col, ok := db.cols.Load(colName)
	if ok {
		if v, ok := col.(*Col); ok {
			return v
		}
	}
	return nil
}

func (db *mongoDbDatabase) SetColByName(colName string) error {
	col := NewMongoColWithDb(colName, db.db)
	if colName != "" {
		db.cols.Store(colName, col)
		return nil
	}
	return errors.New("colName can not null")

}

func (db *mongoDbDatabase) GetMongoDb() *mongo.Database {
	return db.db
}

func (db *mongoDbDatabase) DropAllDatabase() error {
	if err := db.db.Drop(context.Background()); err != nil {
		return nil
	}
	return nil
}

func GetMongoDb(dbName string, opts ...DbOption) (db *mongo.Database) {
	if dbName == "" {
		dbName = viper.GetString("mongo.db")
	}
	if dbName == "" {
		dbName = "test"
	}

	_opts := &DbOptions{}
	for _, op := range opts {
		op(_opts)
	}

	var c *mongo.Client
	if _opts.client == nil {
		var err error
		c, err = GetMongoClient()
		if err != nil {
			trace.PrintError(err)
			return nil
		}
	} else {
		c = _opts.client
	}

	return c.Database(dbName, nil)
}
