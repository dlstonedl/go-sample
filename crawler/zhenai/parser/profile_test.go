package parser

import (
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/dlstonedl/go-sample/crawler/model"
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
	parseResult := ParseProfile(contents, "http://album.zhenai.com/u/1129992868", "Name")

	item := engine.Item{
		Url:  "http://album.zhenai.com/u/1129992868",
		Id:   "1129992868",
		Type: "zhenai",
		Data: model.Profile{
			Name:      "Name",
			City:      "阿坝",
			Gender:    "男",
			Age:       "20",
			Height:    "177cm",
			Weight:    "60kg",
			Income:    "3千以下",
			Marriage:  "未婚",
			Education: "高中及以下",
			Nation:    "汉族",
		},
	}

	if len(parseResult.Items) != 1 || item != parseResult.Items[0] {
		t.Errorf("parseResult.Items: len=%d, item=%v, except len=%d, item=%v",
			len(parseResult.Items), parseResult.Items[0], 1, item)
	}

}
