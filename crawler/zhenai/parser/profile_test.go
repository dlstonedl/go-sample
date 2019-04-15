package parser

import (
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
	parseResult := ParseProfile(contents, "Name", "http://album.zhenai.com/u/1129992868")

	profile := model.Profile{
		Id:        "1129992868",
		Url:       "http://album.zhenai.com/u/1129992868",
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
	}

	if len(parseResult.Items) != 1 || profile != parseResult.Items[0] {
		t.Errorf("parseResult.Items: len=%d, item=%v, except len=%d, item=%v",
			len(parseResult.Items), parseResult.Items[0], 1, profile)
	}

}
