package persist

import (
	"context"
	"fmt"
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/olivere/elastic"
)

type GetEsClientFunc func() *elastic.Client

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
