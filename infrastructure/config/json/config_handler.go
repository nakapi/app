package json

import (
	"app/interface/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
)

type ConfigHandler struct {
	Database DatabaseConfig `json:"database"`
	Logger   LoggerConfig   `json:"logger"`
}

func NewConfigHandler() config.ConfigHandler {
	configHandler := new(ConfigHandler)
	return configHandler
}

func (handler *ConfigHandler) Load() error {
	u, err := user.Current()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	p := filepath.Join(u.HomeDir, "go", "src", "app", "config", "config.json")
	raw, err := ioutil.ReadFile(p)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	json.Unmarshal(raw, &handler)
	return nil
}

func (handler *ConfigHandler) GetDatabaseUser() string {
	return handler.Database.User
}
func (handler *ConfigHandler) GetDatabasePassword() string {
	return handler.Database.Password
}
func (handler *ConfigHandler) GetDatabaseHost() string {
	return handler.Database.Host
}
func (handler *ConfigHandler) GetDatabasePort() string {
	return handler.Database.Port
}
func (handler *ConfigHandler) GetDatabase() string {
	return handler.Database.Database
}

func (handler *ConfigHandler) GetLogger() []byte {
	loggerJson, _ := json.Marshal(handler.Logger)
	return loggerJson
}
