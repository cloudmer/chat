package websocket

import (
	"chat/config"
	"chat/tool"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
)

func (service *service) startWebsocket() error {
	tool.Logger.Info().Str("bind", config.Option.Websocket.Bind).Int("process id", os.Getpid()).Msg("Starting Websocket")
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		service.ws(w, r)
	})
	err := http.ListenAndServe(config.Option.Websocket.Bind, nil)
	return err
}

func (service *service) ws(w http.ResponseWriter, r *http.Request)  {
	up := websocket.Upgrader{
		ReadBufferSize:  config.Option.Websocket.ReadBufferSize,
		WriteBufferSize: config.Option.Websocket.WriteBufferSize,
	}
	// cross origin domain support
	up.CheckOrigin = func(r *http.Request) bool { return true }
	_, err := up.Upgrade(w, r, nil)
	if err != nil {
		tool.Logger.Fatal().Err(err).Msg("websocket 启动失败")
	}
}