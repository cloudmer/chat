package websocket

import (
	"chat/config"
	"chat/tool"
	"runtime"
)

type service struct {

}

func New() *service {
	return new(service)
}

func (service *service) Run()  {
	runtime.GOMAXPROCS(runtime.NumCPU())
	service.readConfig()
	if err := service.startWebsocket(); err != nil {
		tool.Logger.Fatal().Err(err).Msg("Start Websocket fail")
	}
}

// 读取验证配置
func (service *service) readConfig()  {
	if config.Option.Websocket.ReadBufferSize == 0 {
		tool.Logger.Fatal().Msg("请先配置 websocket ReadBufferSize")
	}
	if config.Option.Websocket.WriteBufferSize == 0 {
		tool.Logger.Fatal().Msg("请先配置 websocket WriteBufferSize")
	}
}