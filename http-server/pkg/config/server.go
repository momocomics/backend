package config

import (
	"crypto/rsa"
	"log"

	"github.com/momocomics/backend/http-server/pkg/pb"
)

type ServerConfig struct {
	rpcClient  pb.TodoClient
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
	domain     string
	debug      bool
}

func New(rpcClient pb.TodoClient, privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey, domain string) *ServerConfig {
	return &ServerConfig{rpcClient: rpcClient, privateKey: privateKey, publicKey: publicKey, domain: domain}
}

func (c *ServerConfig) IsDebug() bool {
	return c.debug
}

func (c *ServerConfig) SetDebug(debug bool) {
	c.debug = debug
}

func (c *ServerConfig) RpcClient() pb.TodoClient {
	if c.rpcClient == nil {
		log.Fatal("config: RPC Client not initialised")
	}

	return c.rpcClient
}

func (c *ServerConfig) Domain() string {
	return c.domain
}
