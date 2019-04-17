package persist

import (
	"context"
	"fmt"
	"github.com/dlstonedl/go-sample/crawler/engine"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		successCount := 0
		failCount := 0
		for {
			item := <-out
			log.Printf("ItemSaver item #%d, %v\n", itemCount, item)
			itemCount++

			err := saveTo(client, index, item)
			if err != nil {
				log.Printf("faile save #%d, %v, %v\n", failCount, item, err)
				failCount++
				continue
			}

			log.Printf("success save #%d\n", successCount)
			successCount++
		}
	}()

	return out, nil
}

func saveTo(client *elastic.Client, index string, item engine.Item) error {
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
