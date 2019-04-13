package parser

import (
	"github.com/dlstonedl/go-sample/crawler/engine"
	"regexp"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(contents []byte) engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		city := string(m[2])
		result.Items = append(result.Items, city)
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParseFunc: func(contents []byte) engine.ParseResult {
				return ParseCity(contents, city)
			},
		})
	}
	return result
}
