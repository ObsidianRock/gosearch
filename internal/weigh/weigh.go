package weigh

import (
	"math"

	"github.com/ObsidianRock/gosearch/internal/common"
)

const earthRadius = 6371e3

func toRadians(deg float64) float64 { return deg * math.Pi / 180 }

// DistanceTo return the distance in meteres between two point.
func DistanceTo(latA, lonA, latB, lonB float64) int {
	φ1 := toRadians(latA)
	λ1 := toRadians(lonA)
	φ2 := toRadians(latB)
	λ2 := toRadians(lonB)
	Δφ := φ2 - φ1
	Δλ := λ2 - λ1
	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) + math.Cos(φ1)*math.Cos(φ2)*math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return int((earthRadius * c) + 0.5)
}

// RankScore gives a rank based on number of matching terms in the item name and tokens
func RankScore(itemName string, tokens []string) int {

	var rank int
	tokenize := common.Tokenizer(itemName)
	for _, v := range tokenize {
		if stringInSlice(v, tokens) {
			rank++
		}
	}

	return rank
}

func stringInSlice(s string, arr []string) bool {

	for _, b := range arr {
		if b == s {
			return true
		}
	}
	return false
}
