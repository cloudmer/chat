package websocket

import (
	"chat/config"
	"chat/tool"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
)

func (service *Service) startWebsocket() error {
	tool.Logger.Info().Str("bind", config.Config.Websocket.Bind).Int("process id", os.Getpid()).Msg("Starting Websocket")
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		service.ws(w, r)
	})
	err := http.ListenAndServe(config.Config.Websocket.Bind, nil)
	return err
}

func (service *Service) ws(w http.ResponseWriter, r *http.Request)  {
	up := websocket.Upgrader{
		ReadBufferSize:  config.Config.Websocket.ReadBufferSize,
		WriteBufferSize: config.Config.Websocket.WriteBufferSize,
	}
	// cross origin domain support
	up.CheckOrigin = func(r *http.Request) bool { return true }
	_, err := up.Upgrade(w, r, nil)
	if err != nil {
		tool.Logger.Fatal().Err(err).Msg("websocket 启动失败")
	}
}