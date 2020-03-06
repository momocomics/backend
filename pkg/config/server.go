package config

import (
	"crypto/rsa"

	"github.com/momocomics/backend/pkg/storage"
)

type ServerConfig struct {
	db         storage.Database
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	domain     string
	debug      bool
}

func New(db storage.Database, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, domain string) *ServerConfig {
	return &ServerConfig{db: db, privateKey: privateKey, publicKey: publicKey, domain: domain}
}

func (c *ServerConfig) Domain() string {
	return c.domain
}

func (c *ServerConfig) IsDebug() bool {
	return c.debug
}

func (c *ServerConfig) setDebug(debug bool) {
	c.debug = debug
}
