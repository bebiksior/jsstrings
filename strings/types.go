package jsstrings

type JSStringWithURL struct {
	Value     string     `json:"value"`
	Locations []Location `json:"locations"`
	SourceURL string     `json:"source_url"`
}

type Location struct {
	StartIdx int
	EndIdx   int
}
