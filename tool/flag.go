package tool

import (
	"flag"
	"fmt"
	"os"
)

var start string // 启动服务类型

func ServiceName() string {
	// init logger
	if !LogInitState {
		LoggerInit()
	}

	flag.StringVar(&start, "start", "", "is start service")
	// 改变默认的 Usage
	flag.Usage = Usage
	flag.Parse()

	// select services
	if start == "" {
		Usage()
		Logger.Fatal().Msg("清选择启动服务类型")
	}
	return start
}

func Usage()  {
	fmt.Fprintf(os.Stderr, `chat version: chat/1.0 (Author: cloudmer, github: http://github.com)
-help  get help
-start service1 or service2
`)
}
