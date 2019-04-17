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
	Index       string
	ClientCount int
}

func (es *ElasticSaver) CreateClientPool() chan *elasticsearch.Client {
	var clients []*elasticsearch.Client
	for i := 0; i < es.ClientCount; i++ {
		client, err := elasticsearch.NewDefaultClient()
		if err != nil {
			panic(err)
		}
		clients = append(clients, client)
	}

	clientChan := make(chan *elasticsearch.Client)
	go func() {
		for {
			for _, client := range clients {
				clientChan <- client
			}
		}
	}()
	return clientChan
}

//side-effect，thread-safe
var itemBeforeCount = 0
var successCount = 0
var failCount = 0

func (es *ElasticSaver) Save(client *elasticsearch.Client, item engine.Item) {
	log.Printf("ItemSaver item: #%d, %v\n", itemBeforeCount, item)
	itemBeforeCount++

	err := saveItem(client, es.Index, item)
	if err != nil {
		log.Printf("fail save #%d, %v\n", failCount, item)
		failCount++
		return
	}

	log.Printf("success save: #%d\n", successCount)
	successCount++
}

func saveItem(client *elasticsearch.Client, index string, item engine.Item) (err error) {
	if item.Type == "" || item.Id == "" {
		return fmt.Errorf("must apply Type and Id")
	}

	bytes, err := json.Marshal(item)
	if err != nil {
		return err
	}

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
