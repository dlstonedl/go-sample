package persist

import (
	"context"
	"fmt"
	"github.com/dlstonedl/go-sample/crawler/config"
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/olivere/elastic"
	"log"
)

type SingleSaver struct {
	EsClient EsClient
}

type EsClient interface {
	Init()
	GetEsClient() *elastic.Client
}

type GetEsClientFunc func() *elastic.Client

func (s *SingleSaver) ItemSaver(item engine.Item) error {
	err := save(s.EsClient.GetEsClient(), config.ElasticIndex, item)
	if err != nil {
		log.Printf("fail save %v\n", item)
		return err
	}

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
