package main

import "fmt"

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

func printGreetingInEnglish(bot EngBot) {
	fmt.Println(bot.greeting())
}

func printGreetingInSpanish(bot SpaBot) {
	fmt.Println(bot.greeting())
}
