package site

import (
	"chat/config"
	"chat/tool"
	"fmt"
	"net/http"
)

type service struct {

}

func New() *service {
	return new(service)
}

func (service *service) Run()  {
	// 读取配置文件
	service.readConfig()
	addr := fmt.Sprintf(":%d", config.Config.Site.Port)
	tool.Logger.Info().Str("addr", addr).Msg("site service starting")
	if err := http.ListenAndServe(addr, http.FileServer(http.Dir("./services/site"))); err != nil {
		tool.Logger.Fatal().Err(err).Msg("site 服务启动失败")
	}
}

// 读取验证配置
func (service *service) readConfig()  {
	if config.Config.Site.Port == 0 {
		tool.Logger.Fatal().Msg("请先配置 site port")
	}
}