package persist

import (
	"context"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})
	go func() {
		itemCount := 0
		for {
			item := <-out
			itemCount++
			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver: err saving tiem %v: %v", item, err)
			}
		}
	}()
	return out
}
func save(item interface{}) (id string, err error) {
	client, err := elastic.NewClient(
		elastic.SetSniff(false))
	if err != nil {
		return "", err
	}
	resp, err := client.Index().
		Index("dating_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())
	if err != nil {
		return "", err
	}
	return resp.Id, err
}
