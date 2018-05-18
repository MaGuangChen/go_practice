package main

// var deckSize int

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards := newDeckFromFile("cardFile")
	// // hand, remainingCards is the deal function return value
	// hand, remainingCards := deal(cards, 5)
	// fmt.Println(cards)
	// // hand.print()
	// // remainingCards.print()
	// fmt.Println(hand.print())
	// fmt.Println(remainingCards.print())
	// fmt.Println("under is cards: ")
	// fmt.Println(cards)

	// blue := color("Blue")
	// fmt.Println(blue.describeColor("is a awesome color"))
	// fmt.Println(cards.toString())
	// cards.saveToFile("cardFile")
	// []byte() is []byte type what we want, (greeting) the value what we have

	// var card string = "Ace of Spades" 相等於下方
	// card := "Ace of Spades"
	// card = fiveDiamonds()

	// fmt.Println(card)
	// deckSize = 52
	// fmt.Println(deckSize)

	// newCard := deck{"Ace of Diamonds", fiveDiamonds()}
	// newCard = append(newCard, "Six of heart")

	// for i, card := range newCard {
	// 	fmt.Println(i, card)
	// }
	// newCard.print()
}

func fiveDiamonds() string {
	return "Five of Diamonds"
}
