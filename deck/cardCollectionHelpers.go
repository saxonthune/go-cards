package deck

import (
	"math/rand"
	"sort"
	"time"
)

func TakeTopN(coll *[]Card, n int) []Card {

	if n > len(*coll) {
		result := make([]Card, len(*coll))
		copy(result, *coll)
		*coll = nil
		return result
	}

	result := make([]Card, n)
	copy(result, (*coll)[0:n])
	*coll = (*coll)[n:]
	return result
}

// TakeByIndex
func TakeByIndex(coll *[]Card, idxs []int) []Card {
	result := make([]Card, len(idxs))
	offset := 0

	sort.Sort(sort.IntSlice(idxs))
	for i, v := range idxs {
		result[i] = (*coll)[v-offset]
		*coll = append((*coll)[:v-offset], (*coll)[v-offset+1:]...)
		offset++
	}

	return result
}

func Shuffle(coll *[]Card) {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(*coll), func(i, j int) {
		(*coll)[i], (*coll)[j] = (*coll)[j], (*coll)[i]
	})
}
