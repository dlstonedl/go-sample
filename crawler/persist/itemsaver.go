package persist

import (
	"context"
	"fmt"
	"github.com/dlstonedl/go-sample/crawler/config"
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/olivere/elastic"
	"log"
)

type EsSave struct {
	EsClient EsClient
}

type EsClient interface {
	Init()
	GetEsClient() *elastic.Client
}

//no thread safe
var itemCount = 0
var successCount = 0
var failCount = 0

func (s *EsSave) ItemSave(item engine.Item) error {
	log.Printf("ItemSaver item #%d, %v\n", itemCount, item)
	itemCount++

	err := save(s.EsClient.GetEsClient(), config.ElasticIndex, item)
	if err != nil {
		log.Printf("faile save #%d, %v, %v\n", failCount, item, err)
		failCount++
		return err
	}

	log.Printf("success save #%d\n", successCount)
	successCount++
	return nil
}

func save(client *elastic.Client, index string, item engine.Item) error {
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
