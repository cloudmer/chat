package tool

import (
	"flag"
	"fmt"
	"os"
)

var (
	h     bool // 帮助
	start bool // 启动
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&start, "start", false, "is start service")

	// 改变默认的 Usage
	flag.Usage = usage
	flag.Parse()

	if h {
		usage()
	}
}

func usage()  {
	fmt.Fprintf(os.Stderr, `chat version: chat/1.0
Usage: nginx [-h] [-help] [-start start]
`)
}
