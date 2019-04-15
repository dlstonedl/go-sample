package engine

type ParseFunc func(contents []byte, url string) ParseResult

type Request struct {
	Url       string
	ParseFunc ParseFunc
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
