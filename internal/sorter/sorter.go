package sorter

import (
	"search/storage"
	"sort"
)

type lessFunc func(p1, p2 *storage.Result) bool

type multiSorter struct {
	results []storage.Result
	less    []lessFunc
}

func (ms *multiSorter) Sort(results []storage.Result) {
	ms.results = results
	sort.Sort(ms)
}

func orderedBy(less ...lessFunc) *multiSorter {
	return &multiSorter{
		less: less,
	}
}

func (ms *multiSorter) Len() int {
	return len(ms.results)
}

func (ms *multiSorter) Swap(i, j int) {
	ms.results[i], ms.results[j] = ms.results[j], ms.results[i]
}

func (ms *multiSorter) Less(i, j int) bool {

	p, q := &ms.results[i], &ms.results[j]

	var k int

	for k = 0; k < len(ms.less)-1; k++ {
		less := ms.less[k]
		switch {
		case less(p, q):
			return true
		case less(q, p):
			return false
		}
	}

	return ms.less[k](p, q)
}

//SortResults will sort based on decreasing rank and increasing distance
func SortResults(res []storage.Result) {

	increasingDistance := func(c1, c2 *storage.Result) bool {
		return c1.Distance < c2.Distance
	}

	decreasingRank := func(c1, c2 *storage.Result) bool {
		return c1.Rank > c2.Rank
	}

	orderedBy(decreasingRank, increasingDistance).Sort(res)
}
