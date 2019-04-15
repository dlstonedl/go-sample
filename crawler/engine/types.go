package engine

type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	Url  string
	Id   string
	Type string
	Data interface{}
}

func NilParse([]byte) ParseResult {
	return ParseResult{}
}
