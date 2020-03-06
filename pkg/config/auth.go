package config

import (
	"crypto/rsa"
	"log"
)

func (c *ServerConfig) PrivateKey() *rsa.PrivateKey {

	if c.privateKey == nil {
		log.Fatal("config: private key used to sign JWT token is null")
	}

	return c.privateKey
}

func (c *ServerConfig) PublicKey() *rsa.PublicKey {

	if c.publicKey == nil {
		log.Fatal("config: private key used to sign JWT token is null")
	}

	return c.publicKey
}
