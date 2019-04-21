package persist

import (
	"github.com/dlstonedl/go-sample/crawler/engine"
	"log"
)

type CurrentSaver struct {
	Index    string
	EsClient GetEsClientFunc
}

func (es *CurrentSaver) ItemSaver(item engine.Item) error {
	err := save(es.EsClient(), es.Index, item)
	if err != nil {
		log.Printf("fail save %v\n", item)
		return err
	}

	return nil
}
