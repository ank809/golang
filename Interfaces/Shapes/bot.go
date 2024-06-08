package main

type EnglishBot struct{}
type HindiBot struct{}

type Bot interface {
	getGreeting() string
}

func (eb EnglishBot) getGreeting() string {
	return "Hello "
}

func (hb HindiBot) getGreeting() string {
	return "Namaste"
}

func printGreeting(b Bot) string {
	return b.getGreeting()
}
