package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config interface
type Config interface {
	Load()
}

// ServerConfig Server config obj
type ServerConfig struct {
	Port int
}

// DBConfig database config
type DBConfig struct {
	driver   string
	url      string
	port     int
	dbname   string
	username string
	password string
}

// Load : load config from file
func (c *ServerConfig) Load() {
	filename := "server.json"
	loadFile(filename, c)
}

// Load : load config from file
func (c *DBConfig) Load() {
	filename := "db.json"
	loadFile(filename, c)
}

const _ConfigPath = "\\config\\mbcgo\\"

func loadFile(filename string, cp interface{}) interface{} {

	gopath, found := os.LookupEnv("GOPATH")
	if !found {
		log.Fatal("Setenv GOPATH=<GO Path>")
	}
	fullPath := gopath + _ConfigPath + filename

	log.Println(fullPath)
	file, _ := os.Open(fullPath)

	decoder := json.NewDecoder(file)
	err := decoder.Decode(cp)
	if err != nil {
		log.Println(err)
	}

	return cp
}
