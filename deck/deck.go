package deck

type Deck interface {
	DealCards(count int) []Card
}

type StandardDeck struct {
	Undealt []Card
	Hands   [][]Card
	Piles   [][]Card
}

// InitDeck initializes the deck
func (d *StandardDeck) InitDeck() {

	// todo: reset hands, discard, etc

	d.Undealt = d.DeckDefinition()
}

func (d *StandardDeck) DealHands(numHands int, cardsPerHand int) {
	d.Hands = make([][]Card, numHands)
	for i := range d.Hands {
		d.Hands[i] = d.DealCards(cardsPerHand)
	}
}

func (d *StandardDeck) Shuffle() {
	Shuffle(&d.Undealt)
}

func (d *StandardDeck) DealCards(count int) []Card {

	return TakeTopN(&d.Undealt, count)
	//return CardCollection{}
}

func (d *StandardDeck) MoveCards(dst []Card, src []Card, idxs []int) {
	dst = append(dst, TakeByIndex(&src, idxs)...)
}

func (d *StandardDeck) DeckDefinition() []Card {

	unshuffledDeck := make([]Card, 0, 52)
	suits := []Suit{HEART, DIAMOND, CLUB, SPADE}
	ranks := []Rank{ACE, TWO, THREE, FOUR, FIVE, SIX, SEVEN, EIGHT, NINE, TEN, JACK, QUEEN, KING}

	for _, suit := range suits {
		for _, rank := range ranks {
			value := d.ValorizationFaceCardIsTen(rank)
			card := Card{rank, suit, value}
			unshuffledDeck = append(unshuffledDeck, card)
		}
	}

	return unshuffledDeck
}

func (d *StandardDeck) ValorizationFaceCardIsTen(rank Rank) int {
	if rank >= TEN {
		return 10
	}
	return rank.Pips()
}
