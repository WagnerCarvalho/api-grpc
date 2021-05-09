package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type Config interface {
	Dns() string
	DbName() string
}

type confif struct {
	dbUser string
	dbPass string
	dbHost string
	dbPort int
	dbName string
	dsn    string
}

func NewConfig() Config {
	var err error
	var cfg confif
	cfg.dbUser = os.Getenv("DATABASE_USER")
	cfg.dbPass = os.Getenv("DATABASE_PASS")
	cfg.dbHost = os.Getenv("DATABASE_HOST")
	cfg.dbName = os.Getenv("DATABASE_NAME")
	cfg.dbPort, err = strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		log.Fatalln("Error on load env var:", err.Error())
	}

	cfg.dsn = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", cfg.dbUser, cfg.dbPass, cfg.dbHost, cfg.dbPort, cfg.dbName)
	return &cfg
}

func (c *confif) Dns() string {
	return c.dsn
}

func (c *confif) DbName() string {
	return c.dbName
}
