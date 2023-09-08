package main

func main() {
	// var card string = "Ace of Spades"
	// card := "Ace of Spades"
	// card := newCard()
	// card = "Five of Diamonds"
	// cards := []string{"Ace of Diamonds", newCard()}
	// cards := deck{"Ace of Diamonds", newCard()} //change type to deck
	// cards = append(cards, "Six of Spades")
	cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	//cards := newDeckFromFile("my_cards")
	cards.shuffle()
	cards.print()
	// hand.print()
	// remainingCards.print()
	// fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	// cards.print()
	// for i, card := range cards {
	// 	fmt.Println(i,card)
	// }
	// fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
