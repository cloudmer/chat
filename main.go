package main

import (
	"chat/service"
	"chat/tool"
	_ "chat/tool"
	"fmt"
	"path/filepath"
)

func main() {
	service.Start()
	dir, err := filepath.Abs("./")
	fmt.Println(dir, err)
	tool.Logger.Info().Msg("这是一个INFO 信息")
}
