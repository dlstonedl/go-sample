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

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elasticsearch.NewDefaultClient()
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

func saveTo(client *elasticsearch.Client, index string, item engine.Item) (err error) {
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
