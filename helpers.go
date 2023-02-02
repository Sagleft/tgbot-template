package main

import (
	"errors"
	"io/ioutil"

	tb "gopkg.in/telebot.v3"
)

func readFile(filepath string) ([]byte, error) {
	fileBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.New("failed to read file: " + err.Error())
	}
	return fileBytes, err
}

func tgMessageFilter(upd *tb.Update) bool {
	if upd.Message == nil {
		return true // ignore empty messages
	}

	if upd.Message.Sender.IsBot {
		return false // ignore bots
	}

	if upd.Message.IsService() {
		return false // ignore service messages
	}

	return true
}

func getTgPoller() *tb.MiddlewarePoller {
	poller := &tb.LongPoller{Timeout: longPollerInterval}
	return tb.NewMiddlewarePoller(poller, tgMessageFilter)
}
