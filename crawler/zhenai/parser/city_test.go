package parser

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseCity(t *testing.T) {
	file, err := os.Open("city_test.html")
	if err != nil {
		panic(err)
	}

	contents, err := ioutil.ReadAll(file)
	parseResult := ParseCity(contents)

	var items = []string{
		"静听雨声", "那些感觉", "Trytosmil",
	}

	if len(parseResult.Requests) != 20 {
		t.Errorf("parseResult.Requests len is %d ,expect %d",
			len(parseResult.Requests), 10)
	}

	for i := range items {
		if parseResult.Items[i] != items[i] {
			t.Errorf("parseResult item is %s ,expect %s",
				parseResult.Items[i], items[i])
		}
	}

	if len(parseResult.Items) != 20 {
		t.Errorf("parseResult.Items len is %d ,expect %d",
			len(parseResult.Requests), 10)
	}
}
