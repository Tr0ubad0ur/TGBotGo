package main

import (
	tgadapter "tgbot/adapter"
	"tgbot/config"
	"tgbot/service"
)

func main() {
	conf := config.GetConfig()
	adapter:=tgadapter.NewAdapter(*conf)
	service:=tgservice.NewService(adapter)
	service.Serve()

}
