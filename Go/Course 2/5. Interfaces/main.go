package main

import (
	bot "app/modules/bot"
	shape "app/modules/shape"
)

func main() {
	//bots
	englishBot := bot.EnglishBot{}
	spanishBot := bot.SpanishBot{}

	bot.PrintGreeting(englishBot)
	bot.PrintGreeting(spanishBot)

	//shapes
	triangle := shape.Triangle{Height: 10, Base: 10}
	square := shape.Square{SideLength: 10}

	shape.PrintArea(triangle)
	shape.PrintArea(square)
}
