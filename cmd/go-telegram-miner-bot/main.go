package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/brunoshiroma/go-telegram-miner-bot/internal"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	var minersConfigFileName string

	if len(os.Args) > 1 {
		minersConfigFileName = os.Args[1]
	} else {
		minersConfigFileName = "miners_bot.json" // the default file name
	}

	var minersConfig internal.MinersConfigJson

	minersConfigFile, err := os.Open(minersConfigFileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("%s file not exists, pass existing file name", minersConfigFileName)
			os.Exit(1)
		}
		log.Panic(err)
	}

	buffer := make([]byte, 1024*32)

	read, err := minersConfigFile.Read(buffer)
	if err != nil {
		log.Panic(err)
	}

	err = json.Unmarshal(buffer[0:read], &minersConfig)
	if err != nil {
		log.Panic(err)
	}
	buffer = nil

	username := minersConfig.TelegramUsername
	telegramToken := minersConfig.TelegramToken

	bot, err := tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	autoCheck, err := internal.NewAutoCheck(minersConfig)
	if err != nil {
		log.Printf("Error on initialize autocheck %s", err.Error())
	} else if autoCheck != nil {
		go autoCheck.Start()
	}

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		if update.Message.From.UserName == username { // check if the user is the configured

			// make the requests
			result, err := internal.DoMinersCheck(minersConfig)

			if err != nil {
				log.Printf("Error checking miners, e = %#v", err)
			}

			command := update.Message.Command()
			var msgText string

			switch command {
			case "info":
				msgText = formatInfo(&result)
			case "details":
				msgText = formatDetails(&result)
			default:
				msgText = fmt.Sprintf("%s not implemented", command)
			}

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msgText)
			msg.ReplyToMessageID = update.Message.MessageID
			_, err = bot.Send(msg)
			if err != nil {
				log.Printf("Error sending message, e = %#v", err)
			}
		} else {
			log.Printf("[%s] not allowed", update.Message.From.UserName)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "403")
			msg.ReplyToMessageID = update.Message.MessageID
			_, err = bot.Send(msg)
			if err != nil {
				log.Printf("Error sending message, e = %#v", err)
			}
		}
	}

}

func formatInfo(result *internal.MinersResult) (formated string) {
	for _, miner := range result.Miners {
		formated += fmt.Sprintf("%s = %v\n", miner.Name, miner.Success)
	}
	return
}

func formatDetails(result *internal.MinersResult) (formated string) {
	for _, miner := range result.Miners {
		formated += fmt.Sprintf("%s = %v\n%s\n\n", miner.Name, miner.Success, miner.Result)
	}
	return
}
