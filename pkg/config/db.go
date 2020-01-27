package config

import (
	"log"

	"github.com/momocomics/backend/pkg/storage"
)

func (c *ServerConfig) Db() storage.Database {

	if c.db == nil {
		log.Fatal("config: database not initialised")
	}

	return c.db

}
