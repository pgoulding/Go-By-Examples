package main

func main() {

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()
}

func newCard() string {
	return "5 of Diamonds"
}
