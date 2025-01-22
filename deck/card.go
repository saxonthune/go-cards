package deck

import "strings"

type Card struct {
	Rank  Rank
	Suit  Suit
	Value int
}

// NewCard creates a card by parsing an abbreviated string
// (e.g., 'KS' makes King of Spades, 'TH' makes Ten of Hearts)
// and assigns a value to the card
func NewCard(abbr string) Card {
	if len(abbr) < 2 {
		return Card{}
	}
	return Card{
		Rank:  parseRank(string(abbr[0])),
		Suit:  parseSuit(string(abbr[1])),
		Value: 0,
	}
}

func (c Card) Card(rank Rank, suit Suit) Card {
	return Card{
		Rank: rank,
		Suit: suit,
	}
}

func Equal(c1, c2 Card) bool {
	if c1.Suit == c2.Suit && c1.Value == c2.Value {
		return true
	}
	return false
}

func (c Card) String() string {
	//return fmt.Sprintf("%v of %vs", c.Rank, c.Suit)
	return c.Rank.String()[0:1] + c.Suit.String()[0:1]
}

// Suit Represents the suit of the card.
// Canonical order is: hearts, diamonds, clubs, spades.
type Suit int

const (
	HEART Suit = iota
	DIAMOND
	CLUB
	SPADE
)

func (s Suit) String() string {
	switch s {
	case HEART:
		return "HEART"
	case DIAMOND:
		return "DIAMOND"
	case CLUB:
		return "CLUB"
	case SPADE:
		return `SPADE`
	}
	return ""
}

// Rank represents the rank of the card (ace, two, three, ... king).
type Rank int

const (
	ACE Rank = iota
	TWO
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
)

func (r Rank) String() string {
	switch r {
	case ACE:
		return "A"
	case TWO:
		return "2"
	case THREE:
		return "3"
	case FOUR:
		return "4"
	case FIVE:
		return "5"
	case SIX:
		return "6"
	case SEVEN:
		return "7"
	case EIGHT:
		return "8"
	case NINE:
		return "9"
	case TEN:
		return "T"
	case JACK:
		return "J"
	case QUEEN:
		return "Q"
	case KING:
		return "K"
	}
	return ""
}

func (r Rank) Pips() int {
	switch r {
	case ACE:
		return 1
	case TWO:
		return 2
	case THREE:
		return 3
	case FOUR:
		return 4
	case FIVE:
		return 5
	case SIX:
		return 6
	case SEVEN:
		return 7
	case EIGHT:
		return 8
	case NINE:
		return 9
	case TEN:
		return 10
	case JACK:
		return 11
	case QUEEN:
		return 12
	case KING:
		return 13
	}
	return -1
}

func (c Card) FaceCardHasValue10() int {
	if c.Rank >= TEN {
		return 10
	}
	return c.Rank.Pips()
}

func parseSuit(char string) Suit {
	switch strings.ToUpper(char) {
	case "H":
		return HEART
	case "D":
		return DIAMOND
	case "C":
		return CLUB
	case "S":
		return SPADE
	}
	return -1
}

func parseRank(char string) Rank {
	switch strings.ToUpper(char) {
	case "1":
	case "A":
		return ACE
	case "2":
		return TWO
	case "3":
		return THREE
	case "4":
		return FOUR
	case "5":
		return FIVE
	case "6":
		return SIX
	case "7":
		return SEVEN
	case "8":
		return EIGHT
	case "9":
		return NINE
	case "T":
	case "10":
		return TEN
	case "J":
		return JACK
	case "Q":
		return QUEEN
	case "K":
		return KING
	}
	return -1
}
