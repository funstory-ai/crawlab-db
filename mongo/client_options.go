package mongo

import (
	"context"
	"fmt"
)

type ClientOption func(options *ClientOptions)

type ClientOptions struct {
	Context                 context.Context
	Uri                     string
	Host                    string
	Port                    string
	Db                      string
	Hosts                   []string
	Username                string
	Password                string
	AuthSource              string
	AuthMechanism           string
	AuthMechanismProperties map[string]string
}

func NewMongoClientOption(host, port, db string, opts ...ClientOption) *ClientOptions {
	// client options
	option := &ClientOptions{
		Host:       host,
		Port:       port,
		Db:         db,
		AuthSource: db,
	}
	for _, op := range opts {
		op(option)
	}

	if option.Host == "" {
		option.Host = "localhost"
	}
	if option.Port == "" {
		option.Port = "27017"
	}
	if option.Db == "" {
		option.Db = "crawlab_db"
	}

	if option.AuthSource == "" {
		option.AuthSource = "crawlab_db"
	}
	if option.Uri == "" {
		option.Uri = fmt.Sprintf("mongodb://%s:%s/%s", option.Host, option.Port, option.Db)
	}
	return option
}

func WithContext(ctx context.Context) ClientOption {
	return func(options *ClientOptions) {
		options.Context = ctx
	}
}

func WithUri(value string) ClientOption {
	return func(options *ClientOptions) {
		options.Uri = value
	}
}

func WithHost(value string) ClientOption {
	return func(options *ClientOptions) {
		options.Host = value
	}
}

func WithPort(value string) ClientOption {
	return func(options *ClientOptions) {
		options.Port = value
	}
}

func WithDb(value string) ClientOption {
	return func(options *ClientOptions) {
		options.Db = value
	}
}

func WithHosts(value []string) ClientOption {
	return func(options *ClientOptions) {
		options.Hosts = value
	}
}

func WithUsername(value string) ClientOption {
	return func(options *ClientOptions) {
		options.Username = value
	}
}

func WithPassword(value string) ClientOption {
	return func(options *ClientOptions) {
		options.Password = value
	}
}

func WithAuthSource(value string) ClientOption {
	return func(options *ClientOptions) {
		options.AuthSource = value
	}
}

func WithAuthMechanism(value string) ClientOption {
	return func(options *ClientOptions) {
		options.AuthMechanism = value
	}
}
