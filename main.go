package main

import (
	"chat/services/api"
	"chat/services/site"
	"chat/services/task"
	"chat/services/websocket"
	"chat/tool"
	_ "chat/tool"
	_ "chat/config"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dir, err := filepath.Abs("./")
	fmt.Println(dir, err)
	os.Getpid()
	switch tool.ServiceName() {
	case "websocket":
		websocket.New().Run()
	case "task":
		task.New().Run()
	case "api":
		api.New().Run()
	case "site":
		site.New().Run()
	default:
		tool.Usage()
		tool.Logger.Fatal().Msg("清选择正确的服务类型")
	}
}
