package parser

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseCityList(t *testing.T) {
	file, err := os.Open("citylist_test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	contents, err := ioutil.ReadAll(file)
	parseResult := ParseCityList(contents)

	var requests = [3]string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	const maxSize = 470
	if len(parseResult.Requests) != maxSize {
		t.Errorf("parseResult.Requests len is %d, expect %d\n",
			len(parseResult.Requests), maxSize)
	}

	for i := range requests {
		if parseResult.Requests[i].Url != requests[i] {
			t.Errorf("request is %s, except %s\n",
				parseResult.Requests[i].Url, requests[i])
		}
	}
}
