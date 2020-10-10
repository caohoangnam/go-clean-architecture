package search

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/olivere/elastic/v7"
	"github.com/working/go-clean-architecture/domain"
)

type ElasticRepository struct {
	client *elastic.Client
}

func NewElastic(esURL *string) (*ElasticRepository, error) {

	client, err := elastic.NewClient(
		elastic.SetURL(*esURL),
		elastic.SetSniff(false),
	)
	if err != nil {
		return nil, err
	}
	return &ElasticRepository{client}, nil
}

func (e *ElasticRepository) Close() {}

func (e *ElasticRepository) Create(ctx context.Context, meow domain.Meow) error {
	_, err := e.client.Index().
		Index("meows").
		Type("meow").
		Id(meow.Id).
		BodyJson(meow).
		Refresh("wait_for").
		Do(ctx)
	return err

}

func (e *ElasticRepository) Search(ctx context.Context, query string, skip int64, take int64) ([]domain.Meow, error) {
	result, err := e.client.Search().
		Index("meows").
		Query(
			elastic.NewMultiMatchQuery(query, "body").
				Fuzziness("3").
				PrefixLength(1).
				CutoffFrequency(0.0001),
		).
		//		From(int(skip)).
		//		Size(int(take)).
		Do(ctx)
	fmt.Println("err", err)
	if err != nil {
		return nil, err
	}
	meows := []domain.Meow{}
	for _, hit := range result.Hits.Hits {
		var meow domain.Meow
		if err = json.Unmarshal(hit.Source, &meow); err != nil {
			log.Println(err)
		}
		meows = append(meows, meow)
	}
	return meows, nil
}
