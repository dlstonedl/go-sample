package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/dlstonedl/go-sample/crawler/model"
	"github.com/elastic/go-elasticsearch/v7"
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
			Age:       "20",
			Height:    "177cm",
			Weight:    "60kg",
			Income:    "3千以下",
			Marriage:  "未婚",
			Education: "高中及以下",
			Nation:    "汉族",
		},
	}

	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(err)
	}

	err = save(client, "crawler", expected)
	if err != nil {
		panic(err)
	}

	query := "_id:" + expected.Id
	fmt.Printf("query is %s\n", query)

	response, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithQuery(query),
		client.Search.WithIndex("crawler"),
		client.Search.WithDocumentType(expected.Type),
		client.Search.WithPretty(),
	)
	defer response.Body.Close()

	var r map[string]interface{}
	err = json.NewDecoder(response.Body).Decode(&r)
	if err != nil {
		panic(err)
	}

	source := r["hits"].(map[string]interface{})["hits"].([]interface{})[0].(map[string]interface{})["_source"].(map[string]interface{})

	bytes, err := json.Marshal(source)
	if err != nil {
		panic(err)
	}

	var actual engine.Item
	err = json.Unmarshal(bytes, &actual)
	if err != nil {
		panic(err)
	}

	profile, err := model.GetProfileFromJson(actual.Data)
	actual.Data = profile

	if expected != actual {
		t.Errorf("actual is %+v\n, expect %+v", actual, expected)
	}

}
