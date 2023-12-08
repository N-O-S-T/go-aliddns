//go:build linux || darwin || windows

package main

import (
	"flag"
	"os"
)

var configFile = flag.String("f", os.Getenv("ALIDDNSCONFIG"), "阿里云配置文件")
