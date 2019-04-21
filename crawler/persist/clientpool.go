package persist

import (
	"github.com/dlstonedl/go-sample/crawler/config"
	"github.com/olivere/elastic"
)

type ClientPool struct {
	ClientCount int
	clientChan  chan *elastic.Client
}

func (c *ClientPool) Init() {
	c.ClientCount = config.ClientPoolSize

	var clients []*elastic.Client
	for i := 0; i < c.ClientCount; i++ {
		client, err := elastic.NewClient(elastic.SetSniff(false))
		if err != nil {
			panic(err)
		}
		clients = append(clients, client)
	}

	c.clientChan = make(chan *elastic.Client)
	go func() {
		for {
			for _, client := range clients {
				c.clientChan <- client
			}
		}
	}()
}

func (c *ClientPool) GetEsClient() *elastic.Client {
	return <-c.clientChan
}
