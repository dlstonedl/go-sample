package persist

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"log"
	"strings"
)

type ElasticSaver struct {
	clients     []*elasticsearch.Client
	Index       string
	ClientCount int
}

func (es *ElasticSaver) CreateClientPool() chan *elasticsearch.Client {
	for i := 0; i < es.ClientCount; i++ {
		client, err := elasticsearch.NewDefaultClient()
		if err != nil {
			panic(err)
		}
		es.clients = append(es.clients, client)
	}

	clientChan := make(chan *elasticsearch.Client)
	go func() {
		for {
			for _, client := range es.clients {
				clientChan <- client
			}
		}
	}()
	return clientChan
}

//side-effect，thread-safe
var itemBeforeCount = 0
var itemAfterCount = 0

func (es *ElasticSaver) Save(client *elasticsearch.Client, item engine.Item) {
	log.Printf("ItemSaver item: #%d, %v",
		itemBeforeCount, item)
	itemBeforeCount++
	err := saveItem(client, es.Index, item)
	if err != nil {
		fmt.Errorf("save item Error %v", item)
	}

	log.Printf("after save: #%d, %v",
		itemAfterCount, item)
	itemAfterCount++
}

func saveItem(client *elasticsearch.Client, index string, item engine.Item) (err error) {
	if item.Type == "" || item.Id == "" {
		return fmt.Errorf("must apply Type and Id")
	}

	bytes, err := json.Marshal(item)
	req := esapi.IndexRequest{
		Index:        index,
		DocumentType: item.Type,
		DocumentID:   item.Id,
		Body:         strings.NewReader(string(bytes)),
		Refresh:      "true",
	}

	_, err = req.Do(context.Background(), client)
	return err
}
