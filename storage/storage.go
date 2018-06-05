package storage

type Service interface {
	Search(search string, lat float64, lng float64) ([]Result, error)
	Close() error
}

type Result struct {
	Name     string  `json:"item_name"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
	ItemURL  string  `json:"item_url"`
	ImgURL   string  `json:"img_url"`
	Distance int     `json:"distance"`
	Rank     int     `json:"rank"`
}
