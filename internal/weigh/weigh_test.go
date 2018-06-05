package weigh

import "testing"

func TestRankScore(t *testing.T) {
	tt := []struct {
		name     string
		itemName string
		tokens   []string
		rank     int
	}{
		{"full rank", "Sony A7Sii Kit", []string{"sony", "a7sii", "kit"}, 3},
		{"partial rank", "100-400mm Lens f4.5 L IS II USM", []string{"100400mm", "lens", "f45", "ii"}, 4},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			rank := RankScore(tc.itemName, tc.tokens)

			if rank != tc.rank {
				t.Fatalf("testname: %s, incorrect rank of %v when expecting %v", tc.name, rank, tc.rank)
			}
		})
	}
}
