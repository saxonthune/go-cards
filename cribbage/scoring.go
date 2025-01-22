package cribbage

import "go-learning/deck"

type HandScoreResult struct {
	Score    int
	Hand     []deck.Card
	Starter  deck.Card
	Fifteens [][]deck.Card
	Pairs    [][]deck.Card
	Runs     [][]deck.Card
	hasKnobs bool
}

func ScoreHand(hand []deck.Card, starter deck.Card, isCrib bool) HandScoreResult {
	var score, totalScore int

	score, allFifteens := fifteens(append(hand, starter))
	totalScore += score

	score, allRuns := runs(append(hand, starter))
	totalScore += score

	score, allPairs := pairs(append(hand, starter))
	totalScore += score

	if isCrib {
		if deck.IsFlush(append(hand, starter)) {
			totalScore += len(append(hand, starter))
		}
	} else {
		if deck.IsFlush(append(hand)) {
			totalScore += len(append(hand))
		}
	}

	if hasNobs(append([]deck.Card{starter}, hand...)) {
		totalScore++
	}

	return HandScoreResult{
		Score:    totalScore,
		Hand:     hand,
		Starter:  starter,
		Fifteens: allFifteens,
		Pairs:    allPairs,
		Runs:     allRuns,
	}
}

func fifteens(coll []deck.Card) (score int, matches [][]deck.Card) {
	matches = deck.GetValueMatches(coll, 15)
	return len(matches) * 2, matches
}

func pairs(coll []deck.Card) (score int, matches [][]deck.Card) {
	matches = deck.GetPairs(coll)
	return len(matches) * 2, matches
}

func runs(coll []deck.Card) (score int, matches [][]deck.Card) {
	threes := deck.GetRuns(coll, 3)
	matches = append(matches, threes...)
	score += len(threes) * 3

	fours := deck.GetRuns(coll, 4)
	matches = append(matches, fours...)
	score += len(fours) * 4

	if len(coll) == 5 {
		fives := deck.GetRuns(coll, 5)
		matches = append(matches, fives...)
		score += len(fives) * 5
	}
	return score, matches
}

// checks that first item of
func hasNobs(coll []deck.Card) bool {
	if len(coll) < 2 {
		return false
	}
	for _, v := range coll[1:] {
		if v.Rank == deck.JACK && coll[0].Suit == v.Suit {
			return true
		}
	}
	return false
}
