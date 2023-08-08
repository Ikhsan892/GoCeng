package infrastructures

import (
	"github.com/meilisearch/meilisearch-go"
)

func NewMeilisearch(cfg meilisearch.ClientConfig) *meilisearch.Client {
	return meilisearch.NewClient(cfg)
}
