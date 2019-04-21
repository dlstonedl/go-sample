package persist

import (
	"github.com/dlstonedl/go-sample/crawler/engine"
	"log"
)

type SingleSaver struct {
	Index    string
	EsClient GetEsClientFunc
}

func (s *SingleSaver) ItemSaver(item engine.Item) error {
	err := save(s.EsClient(), s.Index, item)
	if err != nil {
		log.Printf("fail save %v\n", item)
		return err
	}

	return nil
}
