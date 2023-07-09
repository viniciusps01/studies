package bot

import "fmt"

type Bot interface {
	GetGreeting() string
}

func PrintGreeting(bot Bot) {
	fmt.Println(bot.GetGreeting())
}
