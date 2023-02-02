package main

import (
	"fmt"
	"log"

	tb "gopkg.in/telebot.v3"
	"gopkg.in/yaml.v2"
)

func main() {
	configBytes, err := readFile(channelsConfigPath)
	if err != nil {
		log.Fatalf("read config: %v", err)
	}

	var cfg config
	if err := yaml.Unmarshal(configBytes, &cfg); err != nil {
		log.Fatalf("decode config: %v", err)
	}

	b, err := tb.NewBot(tb.Settings{
		Token:  cfg.TelegramToken,
		Poller: getTgPoller(),
	})
	if err != nil {
		log.Fatalf("create tg bot: %v", err)
	}

	b.Handle("/start", func(c tb.Context) error {
		msg := fmt.Sprintf("Hi, your telegram ID is %v", c.Sender().ID)
		return c.Send(msg)
	})

	log.Println("bot started")
	b.Start()
}
