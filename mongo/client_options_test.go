package mongo

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewMongoConnOption(t *testing.T) {
	url := "mongodb://mongodb0.example.com:27017"
	op := NewMongoConnOption(url)
	require.Equal(t, url, op.Url)
}

func TestNewMongoConnOptionWithHost(t *testing.T) {
	url := "mongodb://mongodb0.example.com:27017/test"
	op := NewMongoConnOption("", NewMongoConnOptionWithHost("mongodb0.example.com", "27017", "test"))
	require.Equal(t, "mongodb0.example.com", op.host)
	require.Equal(t, "27017", op.port)
	require.Equal(t, "test", op.db)
	require.Equal(t, url, op.Url)
}

func TestNewMongoClient(t *testing.T) {
	url := "mongodb://localhost:27017/test"
	op := NewMongoConnOption(url)
	_, err := NewMongoClient(op)
	require.Nil(t, err)
}
