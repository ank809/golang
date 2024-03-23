package main

func main() {
	cards := deck{"Ace of diamonds", newcard()}
	cards.print()

}

func newcard() string {
	return "Six of spades"
}
