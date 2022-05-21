package mongo

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/mongo"
	"testing"
)

func TestNewMongoConnOption(t *testing.T) {
	url := "mongodb://mongodb0.example.com:27017"
	op := NewMongoClientOptions(&ClientOptions{Uri: url})
	require.Equal(t, url, op.GetURI())
}

func TestNewMongoClient(t *testing.T) {
	url := "mongodb://localhost:27017/test"
	op := NewMongoClientOptions(&ClientOptions{Uri: url})
	client, err := NewMongoClient(op)
	require.Nil(t, err)
	require.IsType(t, mongo.Client{}, *client.client)
}

func TestClient(t *testing.T) {
	url := "mongodb://localhost:27017/test"
	op := NewMongoClientOptions(&ClientOptions{Uri: url})
	client, err := NewMongoClient(op)
	require.Nil(t, err)
	err = client.Ping()
	fmt.Println(err)
	require.Nil(t, err)
	c := client.GetClient()
	require.IsType(t, c, client.client)
	require.NotNil(t, c)
	dis := client.CloseConn()
	require.Nil(t, dis)
}

func TestClientWithoutUrl(t *testing.T) {
	op := NewMongoClientOptions(&ClientOptions{})
	_, err := NewMongoClient(op)
	require.Nil(t, err)
}

func TestNewClientWithAuth(t *testing.T) {
	op := NewMongoClientOptions(&ClientOptions{Username: "crawlab", Password: "123456"})
	require.Equal(t, "mongodb://localhost:27017/crawlab_db", op.GetURI())
	_, err := NewMongoClient(op)
	require.Nil(t, err)
}
