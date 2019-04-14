package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dlstonedl/go-sample/crawler/model"
	"github.com/elastic/go-elasticsearch/v7"
	"testing"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
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

	id, err := save(expected)
	if err != nil {
		panic(err)
	}

	query := "_id:" + id
	fmt.Printf("query is %s\n", query)

	es, err := elasticsearch.NewDefaultClient()
	response, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithQuery(query),
		es.Search.WithIndex("crawler"),
		es.Search.WithPretty(),
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

	var actual model.Profile
	err = json.Unmarshal(bytes, &actual)
	if err != nil {
		panic(err)
	}

	if expected != actual {
		t.Errorf("actual is %+v\n, expect %+v", actual, expected)
	}

}
