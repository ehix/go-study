package xkcd

type Records struct {
	Records []Record `json:"records"`
}

func (records *Records) AddRecord(r Record) []Record {
	records.Records = append(records.Records, r)
	return records.Records
}

type Record struct {
	URL        string `json:"url"`
	Title      string `json:"safe_title"`
	Day        string `json:"day"`
	Month      string `json:"month"`
	Year       string `json:"year"`
	Transcript string `json:"transcript"`
}
