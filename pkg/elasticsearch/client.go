package elasticsearch

import (
	"github.com/elastic/go-elasticsearch/v7"
)

func NewClient(addresses []string) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: addresses,
	}
	return elasticsearch.NewClient(cfg)
}
