package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	// ErrConfigFileNotFound holds the error when the configuration file is not found.
	ErrConfigFileNotFound = errors.New("configuration file not found")
)

// Config contains the application's configuration.
type Config struct {
	API struct {
		Address        *string       `json:"address"`
		BasePath       *string       `json:"basepath"`
		ReadTimeout    time.Duration `json:"readtimeout"`
		WriteTimeout   time.Duration `json:"writetimeout"`
		MaxHeaderBytes int           `json:"maxheaderbytes"`
		CookieName     *string       `json:"cookiename"`
	} `json:"api"`
	Log struct {
		Level *LogLevel `json:"level"`
		Size  *int      `json:"size"`
		Color *bool     `json:"color"`
	} `json:"log"`
	StateDir string `json:"statedir"`
	Crypto   struct {
		ArgonParams struct {
			Memory      uint32 `json:"memory"`
			Iterations  uint32 `json:"iterations"`
			Parallelism uint8  `json:"parallelism"`
			SaltLength  uint32 `json:"saltlength"`
			KeyLength   uint32 `json:"keylength"`
		} `json:"argonparams"`
	} `json:"crypto"`
}

// ConfigLoad loads the application's configuration from the given file.
func ConfigLoad(configFile string) error {
	var c Config

	defaultLogLevel := Trace
	defaultLogSize := 32
	defaultLogColor := true
	defaultAddress := "127.0.0.1:8088"
	defaultBasePath := "/"
	defaultCookieName := "session"

	c.Log.Level = &defaultLogLevel
	c.Log.Size = &defaultLogSize
	c.Log.Color = &defaultLogColor
	c.API.Address = &defaultAddress
	c.API.BasePath = &defaultBasePath
	c.API.CookieName = &defaultCookieName
	c.API.ReadTimeout = 10 * time.Second
	c.API.WriteTimeout = 10 * time.Second
	c.API.MaxHeaderBytes = 1 << 20
	c.Crypto.ArgonParams.Memory = 8 * 1024
	c.Crypto.ArgonParams.Iterations = 3
	c.Crypto.ArgonParams.Parallelism = 2
	c.Crypto.ArgonParams.SaltLength = 16
	c.Crypto.ArgonParams.KeyLength = 32

	app.config = c

	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Error("ioutil.ReadFile: %v", err)
		return ErrConfigFileNotFound
	}

	if err = yaml.Unmarshal(data, &c); err != nil {
		return fmt.Errorf("yaml.Unmarshal: %v", err)
	}

	log.Trace("loaded configuration from %s:", configFile)

	yaml, err := yaml.Marshal(&c)
	if err != nil {
		log.Error("YAML marshal: %v", err)
	} else {
		for _, line := range strings.Split(strings.Trim(string(yaml), "\n"), "\n") {
			log.Trace("%s", line)
		}
	}

	app.config = c

	return nil
}
