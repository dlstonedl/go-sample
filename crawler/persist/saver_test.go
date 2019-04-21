package persist

import (
	"context"
	"encoding/json"
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/dlstonedl/go-sample/crawler/model"
	"github.com/olivere/elastic"
	"testing"
)

func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:  "http://album.zhenai.com/u/1129992868",
		Type: "zhenai",
		Id:   "1129992868",
		Data: model.Profile{
			Name:      "Name",
			City:      "阿坝",
			Gender:    "男",
			Age:       20,
			Height:    177,
			Weight:    60,
			Income:    "3千以下",
			Marriage:  "未婚",
			Education: "高中及以下",
			Nation:    "汉族",
		},
	}

	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	save(client, "crawler_test", expected)

	response, err := client.Get().
		Index("crawler_test").
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal([]byte(response.Source), &actual)

	if err != nil {
		panic(err)
	}

	profile, err := model.GetProfileFromJson(actual.Data)
	if err != nil {
		panic(err)
	}
	actual.Data = profile

	if expected != actual {
		t.Errorf("actual is %+v\n, expect %+v", actual, expected)
	}
}
