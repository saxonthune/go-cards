package main

import (
	"fmt"
	deck2 "go-learning/deck"
)

func main() {

	//deckTesting()
	//cribbageRoundInitTesting()
	//fifteenTesting()
	//RunTesting()
	FlushTesting()

}

func FlushTesting() {
	hand := []deck2.Card{
		{Rank: deck2.QUEEN, Suit: deck2.HEART, Value: 10},
		{Rank: deck2.JACK, Suit: deck2.HEART, Value: 10},
		{Rank: deck2.KING, Suit: deck2.HEART, Value: 10},
	}
	fmt.Println(deck2.IsFlush(hand))

	hand = append(hand, deck2.Card{Rank: deck2.KING, Suit: deck2.DIAMOND, Value: 10})
	fmt.Println(deck2.IsFlush(hand))
}

func RunTesting() {
	deck := deck2.StandardDeck{}
	deck.InitDeck()
	deck.Shuffle()

	deck.DealHands(1, 8)

	fmt.Println(deck.Hands[0])
	runs := deck2.GetRuns(deck.Hands[0], 3)
	fmt.Println(runs)

	//hand := []deck2.Card{
	//	{Rank: deck2.QUEEN, Suit: deck2.HEART, Value: 10},
	//	{Rank: deck2.JACK, Suit: deck2.HEART, Value: 10},
	//	{Rank: deck2.KING, Suit: deck2.HEART, Value: 10},
	//}
	//fmt.Println(deck2.GetRuns(hand, 3))
}

func fifteenTesting() {
	deck := deck2.StandardDeck{}
	deck.InitDeck()
	deck.Shuffle()

	deck.DealHands(1, 5)

	fmt.Println(deck.Hands[0])
	fifteens := deck2.GetValueMatches(deck.Hands[0], 15)
	fmt.Println(fifteens)
}

func cribbageRoundInitTesting() {

	deck := deck2.StandardDeck{}
	deck.InitDeck()
	deck.Shuffle()

	deck.DealHands(2, 6)
	for i := range deck.Hands {
		fmt.Println(deck.Hands[i])
	}

	deck.Piles = make([][]deck2.Card, 2) // crib, turn card

	deck.Piles[0] = append(deck.Piles[0], deck2.TakeByIndex(&deck.Hands[0], []int{0, 1})...)
	deck.Piles[0] = append(deck.Piles[0], deck2.TakeByIndex(&deck.Hands[1], []int{4, 5})...)

	deck.Piles[1] = deck.DealCards(1)

	fmt.Println("------------")
	fmt.Println(deck.Hands[0])
	fmt.Println(deck.Hands[1])
	fmt.Println("Crib: ", deck.Piles[0])
	fmt.Println("Starter: ", deck.Piles[1])
	fmt.Println(deck.Undealt)

	//fmt.Println(deck.Undealt)
	//engine := cribbage.Engine{}

}

func deckTesting() {
	deck := deck2.StandardDeck{}
	deck.InitDeck()

	//card := deckOfCards.Card{
	//	Rank: deckOfCards.KING,
	//	Suit: deckOfCards.HEART,
	//}
	//fmt.Println(card)
	//
	//firstFive := deck.DealCards(5)
	//fmt.Println(firstFive)

	//fmt.Println(fmt.Sprintf("%v", deck.DealCards(5)))
	//fmt.Println(deck.DealCards(5))

	//deck.InitDeck()
	//for deck.Undealt.Cards != nil {
	//	fmt.Println(deck.DealCards(12))
	//}

	//deck.InitDeck()
	//for deck.Undealt.Cards != nil {
	//	test := deckOfCards.TakeTopN(&deck.Undealt.Cards, 12)
	//	fmt.Println(test)
	//}

	deck.InitDeck()
	deck.Shuffle()
	for deck.Undealt != nil {
		fmt.Println(deck2.TakeTopN(&deck.Undealt, 12))
		fmt.Println(deck.DealCards(5))
	}

	fmt.Println(deck2.TakeTopN(&deck.Undealt, 12))
}
