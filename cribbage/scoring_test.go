package cribbage

import (
	"go-learning/deck"
	"reflect"
	"testing"
)

func TestScoreHand(t *testing.T) {
	type args struct {
		in0 []deck.Card
		in1 deck.Card
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5H 5D 5C JS, 5S should return 29",
			args: args{[]deck.Card{deck.NewCard("5H"), deck.NewCard("5D"), deck.NewCard("5C"), deck.NewCard("JS")},
				deck.NewCard("5S")},
			want: 29,
		},
		{
			name: "JH 4D 8C 9S, QS should return 0",
			args: args{[]deck.Card{deck.NewCard("JH"), deck.NewCard("4D"), deck.NewCard("8C"), deck.NewCard("9S")},
				deck.NewCard("QS")},
			want: 0,
		},
		{
			name: "KS 4S 8S 9S, QS should return 5 (crib flush)",
			args: args{[]deck.Card{deck.NewCard("KS"), deck.NewCard("4S"), deck.NewCard("8S"), deck.NewCard("9S")},
				deck.NewCard("QS")},
			want: 5,
		},
		{
			name: "JS 4D 8C 9S, QS should return 1 (nobs)",
			args: args{[]deck.Card{deck.NewCard("JS"), deck.NewCard("4D"), deck.NewCard("8C"), deck.NewCard("9S")},
				deck.NewCard("QS")},
			want: 1,
		},
		{
			name: "AS 2D 2C 3S, 6S should return 8 (double run)",
			args: args{[]deck.Card{deck.NewCard("AS"), deck.NewCard("2D"), deck.NewCard("2C"), deck.NewCard("3S")},
				deck.NewCard("6S")},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := range tt.args.in0 {
				tt.args.in0[i].Value = tt.args.in0[i].FaceCardHasValue10()
			}
			tt.args.in1.Value = tt.args.in1.FaceCardHasValue10()

			if got := ScoreHand(tt.args.in0, tt.args.in1, true); !reflect.DeepEqual(got.Score, tt.want) {
				t.Errorf("ScoreHand() = %v, want %v", got, tt.want)
			}
		})
	}
}
