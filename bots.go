package main

import "fmt"

// Bot expected behavior for any bot
type Bot interface {
	greeting() string
}

// EngBot bot for english
type EngBot struct {}

func (EngBot) greeting() string {
	return "Hello!!";
}

// SpaBot bot for spanish
type SpaBot struct {}

func (SpaBot) greeting() string {
	return "Hola";
}

func printGreeting(bot Bot) {
	fmt.Println(bot.greeting())
}
