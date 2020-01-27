package config

import "github.com/momocomics/backend/pkg/storage"

type ServerConfig struct {
	db storage.Database
}

func New(db storage.Database) *ServerConfig {
	return &ServerConfig{db: db}
}
