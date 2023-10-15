package main

// var deckSize int

func main() {

	// deckSize = 52
	// var card string = "Ace of Spades"
	// card := newCard()
	// fmt.Println(card)
	// printState()
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// cards.print()

	// greeting := "Hi there!"
	// fmt.Println([]byte(cards[0]))

	// c := color("Red")
	// fmt.Println(c.describe("is an awesome color"))
	// cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}

// type color string

// func (c color) describe(decription string) string {
// 	return string(c) + " " + decription
// }
