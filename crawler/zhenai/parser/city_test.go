package parser

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseCity(t *testing.T) {
	file, err := os.Open("city_test.txt")
	if err != nil {
		panic(err)
	}

	contents, err := ioutil.ReadAll(file)
	parseResult := ParseCity(contents, "")

	if len(parseResult.Requests) != 98 {
		t.Errorf("parseResult.Requests len is %d ,expect %d",
			len(parseResult.Requests), 98)
	}
}
