package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
)

type Config struct {
	Mysql struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		Database string `json:"database"`
	} `json:"mysql"`
}

func Load() Config {
	var config Config
	u, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	p := filepath.Join(u.HomeDir, "go", "src", "app", "infrastructure", "config", "config.json")
	raw, err := ioutil.ReadFile(p)
	json.Unmarshal(raw, &config)
	return config
}
