package main

import (
	"encoding/json"
	"errors"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

type APIConfig struct {
	AuthToken string `json:"auth_token"`
	Host      string `json:"host"`
	EntityId  string `json:"entity_id"`
	Resources struct {
		State   string
		Service string
	}
}

type Config struct {
	API APIConfig
}

func (c *Config) Load(path string) {
	flagPath := path

	if path == "" {
		configDir, err := os.UserConfigDir()

		if err != nil {
			log.Fatal(err)
		}

		path = configDir + "/light-control/config.json"
	}

	jsonFile, err := os.Open(path)

	if err != nil {
		if flagPath == "" && errors.Is(err, fs.ErrNotExist) {
			c.CheckConfig()
			return
		}

		log.Fatal(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	clone := c.API

	json.Unmarshal(byteValue, &c.API)

	if clone.EntityId != "" {
		c.API.EntityId = clone.EntityId
	}

	if clone.AuthToken != "" {
		c.API.AuthToken = clone.AuthToken
	}

	if clone.Host != "" {
		c.API.Host = clone.Host
	}

	c.CheckConfig()
}

func (c *Config) CheckConfig() {
	if c.API.EntityId == "" {
		log.Fatal("No entity id given! Exiting...")
	}

	if c.API.AuthToken == "" {
		log.Fatal("No auth token given! Exiting...")
	}

	if c.API.Host == "" {
		log.Fatal("Host not set! Exiting...")
	}
}
