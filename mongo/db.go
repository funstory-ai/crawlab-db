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

type MongoDbDatabase struct {
	dbName string
	client *mongo.Client // may be connection pool
	db     *mongo.Database
	cols   sync.Map
}

func NewMongoDbDatabase(name string, client *mongo.Client) *MongoDbDatabase {
	if name == "" {
		name = defaultDbName
	}
	db := &MongoDbDatabase{dbName: name, client: client}
	db.db = client.Database(name)
	return db
}

func (db *MongoDbDatabase) GetClient() *mongo.Client {
	return db.client
}

func (db *MongoDbDatabase) GetColByName(colName string) *Col {
	col, ok := db.cols.Load(colName)
	if ok {
		if v, ok := col.(*Col); ok {
			return v
		}
	}
	return nil
}

func (db *MongoDbDatabase) SetColByName(colName string) (*Col, error) {
	if colName == "" {
		return nil, errors.New("colName can not null")
	}
	col, ok := db.cols.Load(colName)
	if ok {
		if v, ok := col.(*Col); ok {
			return v, nil
		}
	}
	v := NewMongoColWithDb(colName, db.db)
	db.cols.Store(colName, v)
	return v, nil

}

func (db *MongoDbDatabase) GetMongoDb() *mongo.Database {
	return db.db
}

func (db *MongoDbDatabase) DropAllDatabase() error {
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
