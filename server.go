package main

import (
	"fmt"
	"mbcgo/config"
)

func main() {
	conf := new(config.ServerConfig)
	conf.Load()
	fmt.Println(conf)
}
