package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	LogicConfig LogicConfig
}

type LogicConfig struct {
	
}

func init()  {
	path, _ := filepath.Abs("./")
	configFilePath := path+"/config/config.yaml1"
	_, err := os.Stat(configFilePath)
	if os.IsNotExist(err) {
		// todo 错误
		panic(err)
	}
}