package persist

import (
	"context"
	"fmt"
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/olivere/elastic"
	"log"
)

type ElasticSaver struct {
	Index       string
	ClientCount int
}

func (es *ElasticSaver) CreateClientPool() chan *elastic.Client {
	var clients []*elastic.Client
	for i := 0; i < es.ClientCount; i++ {
		client, err := elastic.NewClient(elastic.SetSniff(false))
		if err != nil {
			panic(err)
		}
		clients = append(clients, client)
	}

	clientChan := make(chan *elastic.Client)
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

func (es *ElasticSaver) Save(client *elastic.Client, item engine.Item) {
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

func saveItem(client *elastic.Client, index string, item engine.Item) error {
	if item.Type == "" || item.Id == "" {
		return fmt.Errorf("must apply Type and Id")
	}

	_, err := client.Index().
		Index(index).
		Type(item.Type).
		Id(item.Id).
		BodyJson(item).
		Do(context.Background())

	return err
}
