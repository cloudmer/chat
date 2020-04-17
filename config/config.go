package config

import (
	"chat/tool"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
	"sync"
)

const (
	yamlConfigPath string = "/config/config.yaml"
)

var (
	configOnce sync.Once
	Config config
)

type config struct {
	Websocket websocket `yaml:"websocket"`
	Site 	  site 		`yaml:"site"`
}

type websocket struct {
	Bind 			string  `yaml:"bind"`
	ReadBufferSize  int 	`yaml:"ReadBufferSize"`
	WriteBufferSize int 	`yaml:"WriteBufferSize"`
}

type site struct {
	Port int `yaml:"port"`
}

func init()  {
	initConfig()
}

func initConfig()  {
	configOnce.Do(func() {
		// init logger
		if !tool.LogInitState {
			tool.LoggerInit()
		}

		path, _ := filepath.Abs("./")
		configFilePath := path+yamlConfigPath
		yamlConfig, err := ioutil.ReadFile(configFilePath)
		if err != nil {
			// todo 错误
			tool.Logger.Fatal().Err(err).Msg("未找到该文件")
		}
		yaml.Unmarshal(yamlConfig, &Config)
	})
}