package sqlite

import (
	"database/sql"
	"fmt"
	"search/storage"

	"github.com/ObsidianRock/gosearch/internal/common"
	"github.com/ObsidianRock/gosearch/internal/sorter"
	"github.com/ObsidianRock/gosearch/internal/weigh"

	_ "github.com/mattn/go-sqlite3"
)

type sqlite struct {
	db *sql.DB
}

type input struct {
	lat        float64 `json:"lat"`
	lng        float64 `json:"lng"`
	searchTerm string  `json:"searchTerm"`
	tokenized  []string
	query      string
}

// New returns an instance of the SQLite storage which implements the Service interface
func New(path string) (storage.Service, error) {

	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("failed opening db connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping db: %v", err)
	}

	return &sqlite{db}, nil
}

func (s sqlite) Close() error {
	return s.Close()
}

//
func (s sqlite) Search(search string, lat float64, lng float64) ([]storage.Result, error) {
	i := input{searchTerm: search, lat: lat, lng: lng}

	i.tokenized = common.Tokenizer(i.searchTerm)
	i.query = genQuery(i.tokenized)

	// converting []string to []interface for sql query
	qinterface := make([]interface{}, len(i.tokenized))

	for i, v := range i.tokenized {
		qinterface[i] = "%" + v + "%"
	}

	rows, err := s.db.Query(i.query, qinterface...)
	if err != nil {
		return nil, fmt.Errorf("failed making db query: %v", err)
	}

	var results []storage.Result
	for rows.Next() {

		var rslt storage.Result

		err := rows.Scan(&rslt.Name, &rslt.Lat, &rslt.Lng, &rslt.ItemURL, &rslt.ImgURL)
		if err != nil {
			return nil, fmt.Errorf("failed scanning result to struct: %v", err)
		}

		rslt.Distance = weigh.DistanceTo(i.lat, i.lng, rslt.Lat, rslt.Lng)

		rslt.Rank = weigh.RankScore(rslt.Name, i.tokenized)

		results = append(results, rslt)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("failed iterting db rows: %v", err)
	}

	sorter.SortResults(results)

	defer rows.Close()

	if len(results) == 0 {
		results = []storage.Result{}
	} else if len(results) <= 20 {
		results = results[:len(results)-1]
	} else {
		results = results[:20]
	}

	return results, nil
}

func genQuery(tokens []string) string {

	baseq := "SELECT * FROM items WHERE "
	var placeholders string

	for j := 0; j < len(tokens); j++ {

		if j+1 == len(tokens) {
			placeholders += "item_name LIKE ? "
		} else {
			placeholders += "item_name LIKE ? OR "
		}
	}

	return baseq + placeholders
}
