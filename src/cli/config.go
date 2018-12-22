package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

type Config struct {
	Endpoint    string `json:"endpoint"`
	AccessToken string `json:"access_token"`
}

func loadConfig(configPath string, flags *ApplicationFlags) (*Config, error) {
	cfgPath := configPath
	userSpecified := cfgPath!= ""

	if !userSpecified {
		user, err := user.Current()
		if err != nil {
			return nil, err
		}

		cfgPath = filepath.Join(user.HomeDir, "cde-config.json")
	}

	data, err := ioutil.ReadFile(os.ExpandEnv(cfgPath))

	if err != nil && (!os.IsNotExist(err) || userSpecified) {
		return nil, err
	}

	var cfg Config

	if err == nil {
		if err := json.Unmarshal(data, &cfg); err != nil {
			return nil, err
		}
	}

	// Apply config overrides.
	if envToken := os.Getenv("SRC_ACCESS_TOKEN"); envToken != "" {
		cfg.AccessToken = envToken
	}

	// Override endpoint from flags
	if flags.Endpoint != "" {
		cfg.Endpoint = strings.TrimSuffix(flags.Endpoint, "/")
	}

	// Override access_key from flags
	if flags.AccessToken != "" {
		cfg.AccessToken = flags.AccessToken
	}

	return &cfg, nil
}
