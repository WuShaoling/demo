package main

import (
	"github.com/demo/logger"
	"time"
)

func main() {
	logger.Init()
	for {
		logger.Log.Info("hello world")
		time.Sleep(1 * time.Second)
	}

}
