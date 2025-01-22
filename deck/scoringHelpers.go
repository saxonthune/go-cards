package deck

import "sort"

func GetValueMatches(cards []Card, targetValue int) [][]Card {
	var result [][]Card
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].Value < cards[j].Value })
	getValueMatchesDfs(cards, 0, targetValue, []Card{}, &result)
	return result
}

func getValueMatchesDfs(cards []Card, start int, targetSum int, current []Card, result *[][]Card) {
	if targetSum < 0 {
		return
	}

	if targetSum == 0 {
		t := make([]Card, len(current))
		copy(t, current)
		*result = append(*result, t)
		return
	}

	for i := start; i < len(cards); i++ {
		if i > start && cards[i] == cards[i-1] {
			continue
		}

		current = append(current, cards[i])
		getValueMatchesDfs(cards, i+1, targetSum-cards[i].Value, current, result)
		current = current[:len(current)-1]
	}
}

// GetRuns finds all unique runs of length _k_. Assumes that aces are low.
func GetRuns(cards []Card, k int) [][]Card {
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].Rank < cards[j].Rank })
	var result [][]Card
	if len(cards) == k {
		return getFullRun(cards)
	}
	getRunsRecursive(cards, 0, k, []Card{}, &result)
	return result
}

// getRuns checks if the (sorted) slice cards
func getFullRun(cards []Card) [][]Card {
	for i := 1; i < len(cards); i++ {
		if cards[i-1].Rank != cards[i].Rank-1 {
			return [][]Card{}
		}
	}
	result := make([]Card, len(cards))
	copy(result, cards)
	return [][]Card{result}
}

func getRunsRecursive(cards []Card, startIndex int, targetLength int, current []Card, result *[][]Card) {
	if len(current) > 1 {
		if current[len(current)-1].Rank != current[len(current)-2].Rank+1 {
			return
		}
	}

	if targetLength == 0 {
		t := make([]Card, len(current))
		copy(t, current)
		*result = append(*result, t)
		return
	}

	for i := startIndex; i < len(cards); i++ {
		if i > startIndex && cards[i] == cards[i-1] {
			continue
		}

		current = append(current, cards[i])
		getRunsRecursive(cards, i+1, targetLength-1, current, result)
		current = current[:len(current)-1]
	}
}

func IsFlush(cards []Card) bool {
	if len(cards) < 2 {
		return false
	}

	for i := 1; i < len(cards); i++ {
		if cards[i-1].Suit != cards[i].Suit {
			return false
		}
	}
	return true
}

func GetPairs(cards []Card) (matches [][]Card) {
	sort.SliceStable(cards, func(i, j int) bool { return cards[i].Rank < cards[j].Rank })
	for i, vi := range cards {
		for j, vj := range cards[i+1:] {
			if vi.Rank == vj.Rank {
				t := make([]Card, 2)
				copy(t, []Card{cards[i], cards[j]})
				matches = append(matches, t)
			} else {
				break
			}
		}
	}
	return matches
}
