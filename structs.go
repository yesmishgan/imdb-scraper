package imdb_scraper

type Poster struct {
	Height uint16 `json:"height"`
	Width  uint16 `json:"width"`
	URL    string `json:"url"`
}

type Primary struct {
	Href  string `json:"href"`
	Year  uint16
	Title string
}

type Film struct {
	Primary Primary `json:"primary"`
	Plot    string  `json:"plot"`
	Poster  Poster  `json:"poster"`

}
