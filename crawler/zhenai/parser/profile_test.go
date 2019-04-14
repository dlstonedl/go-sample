package parser

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestParseProfile(t *testing.T) {
	file, err := os.Open("profile_test.txt")
	if err != nil {
		panic(err)
	}

	contents, err := ioutil.ReadAll(file)
	ParseProfile(contents, "", "")
}
