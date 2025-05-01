package main

import (
	"log"
	"time"

	tele "gopkg.in/telebot.v4"
)

func main() {
	pref := tele.Settings{
		Token:  "",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hello! I am echo-bot")
	})
	b.Handle(tele.OnText, func(c tele.Context) error {
		return c.Send("You wrote: " + c.Text())
	})
	b.Handle(tele.OnPhoto, func(c tele.Context) error {
		return c.Send(c.Message().Photo)
	})
	b.Start()
}
