package service

import (
	"fmt"
	"os"
)

// service start
func Start()  {
	output()
}

// output msg
func output()  {
	fmt.Fprintf(os.Stderr, `chat version: chat/1.0
Author: cloudmer
github: http://github.com
`)
}
