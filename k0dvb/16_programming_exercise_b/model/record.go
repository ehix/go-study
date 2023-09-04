package ex16b

type Record struct {
	URL        string
	Title      string `json:"safe_title"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Transcript string `json:"transcript"`
}
