package mongo

import (
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestMongoGetDb(t *testing.T) {
	dbName := "test_db"
	viper.Set("mongo.db", dbName)

	db := GetMongoDb("")
	require.Equal(t, dbName, db.Name())
}

func setupClient() *mongo.Client {
	url := "mongodb://localhost:27017/crwalab_test"
	c, err := NewMongoClient(NewMongoClientOptions(&ClientOptions{Uri: url}))
	if err != nil {
		panic(err)
	}
	return c.client

}

func setupDb() *MongoDbDatabase {
	c := setupClient()
	db := NewMongoDbDatabase("crawlab_db", c)
	return db
}

func TestNewMongoDbDatabase(t *testing.T) {
	c := setupClient()
	db := NewMongoDbDatabase("crawlab_db", c)
	require.Equal(t, c, db.client)
	require.IsType(t, mongo.Database{}, *db.db)
	col := db.GetColByName("")
	require.Nil(t, col)
	_, err := db.SetColByName("")
	require.Errorf(t, err, "colName can not null")
	col_1, err := db.SetColByName("tag")
	require.Nil(t, err)
	col_2 := db.GetColByName("tag")
	require.NotNil(t, col_2)
	require.Equal(t, col_1, col_2)
	require.Equal(t, col_2.c.Name(), "tag")
}
