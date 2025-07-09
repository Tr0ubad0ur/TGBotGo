package tgadapter

import (
	"tgbot/config"
	"log"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IncomingMessage struct {
	ChatID int64
	Text string
}

type IncomingCommand struct {
}

type SendTaskListCommand struct {
}

type Adapter struct {
	Config config.BotConfig
	bot *tgbotapi.BotAPI
	messageChan chan IncomingMessage
	commandChan chan IncomingCommand

}

func (a *Adapter) GetMessageChannel() <-chan IncomingMessage {
	return a.messageChan
}

func (a *Adapter) GetCommandChannel() <-chan IncomingCommand {
	return a.commandChan
}

func (a *Adapter) SendTaskList(c SendTaskListCommand) error {
	return nil
}

func NewAdapter(c config.BotConfig) *Adapter {
	a := &Adapter{
		Config: c,
	}

	a.messageChan=make(chan IncomingMessage)
	a.commandChan=make(chan IncomingCommand)

	if err := a.start(); err!= nil{
		panic(err)
	}

	return a
}

func (a *Adapter) start() error {
	var err error
	a.bot, err = tgbotapi.NewBotAPI(a.Config.APIKey)
	if err != nil {
		return err
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = a.Config.Timeout

	updates := a.bot.GetUpdatesChan(u)

	go func()  {
		for upd:=range updates {
			switch {
			case upd.Message.Text!="":
				if upd.Message.Chat.ID != 123456789 {
					log.Println("Foregin user send message")
					continue
				}
				msg:=IncomingMessage{
					ChatID: upd.Message.Chat.ID,
					Text: upd.Message.Text,
				}
				a.messageChan <- msg
			case upd.CallbackQuery!=nil:
				
			}
		}
	}()
	return nil
}
