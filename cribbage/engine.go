package cribbage

import (
	deck2 "go-learning/deck"
)

type Engine struct {
	Deck deck2.StandardDeck
}

// run a hand
func (e Engine) PlayHand() {

	e.Deck.InitDeck()
	e.Deck.Shuffle()

	// deal two hands of 6 cards
	e.Deck.DealHands(2, 6)

	// put two cards from each hand into crib

}
