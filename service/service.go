package tgservice

import (
	"fmt"
	"log"
	"tgbot/adapter"
)

type TgAPIAdapter interface {
	GetMessageChannel() <-chan tgadapter.IncomingMessage
	GetCommandChannel() <-chan tgadapter.IncomingCommand
	SendTaskList(tgadapter.SendTaskListCommand) error
}

type Service struct{
	adapter TgAPIAdapter
}

func NewService(a TgAPIAdapter) *Service{
	srv := &Service{adapter: a}

	
	return srv
}

func (s *Service) Serve() error{

	log.Println("ServeStarted")
	for{
		select{
		case msg:=<-s.adapter.GetMessageChannel():
			fmt.Println(msg.Text)
		}
	}
}
