package persist

import "github.com/olivere/elastic"

type SingleClient struct {
	client *elastic.Client
}

func (s *SingleClient) Init() {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}

	s.client = client
}

func (s *SingleClient) GetEsClient() *elastic.Client {
	return s.client
}
