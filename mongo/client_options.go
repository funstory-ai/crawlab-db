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

type ConnOption struct {
	Url                     string
	base                    string
	host                    string
	port                    string
	db                      string
	Hosts                   []string
	username                string
	password                string
	authSource              string
	authMechanism           string
	authMechanismProperties map[string]string
}

type ConnOptionFunc func(options *ConnOption)

func NewMongoConnOption(url string, opts ...ConnOptionFunc) *ConnOption {
	if url != "" {
		return &ConnOption{Url: url}
	}
	// client options
	option := &ConnOption{
		base: "mongodb://",
	}
	if option.host == "" {
		option.host = "localhost"
	}
	if option.port == "" {
		option.port = "27017"
	}
	if option.db == "" {
		option.db = "crawlab_db"
	}

	if option.authSource == "" {
		option.authSource = "crawlab_db"
	}
	if option.Url == "" {
		option.Url = fmt.Sprintf("mongodb://%s:%s/%s", option.host, option.port, option.db)
	}
	for _, op := range opts {
		op(option)
	}
	return option
}

func NewMongoConnOptionWithAuth(username, password string) ConnOptionFunc {
	return func(options *ConnOption) {
		options.username = username
		options.password = password
		options.Url = fmt.Sprintf("%s%s:%s@%s:%s/%s", options.base, options.username, options.password,
			options.host, options.port, options.db)
	}
}

func NewMongoConnOptionWithHost(host, port, db string) ConnOptionFunc {
	return func(options *ConnOption) {
		options.host = host
		options.port = port
		options.db = db
		options.Url = fmt.Sprintf("%s%s:%s/%s", options.base, options.host, options.port, options.db)
	}
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
