package routes

import (
	"github.com/8rxn/go-shortener/database"
)

type Response struct {
	Slug   string `json:"slug"`
	Url    string `json:"url"`
	Expiry string `json:"expiry,omitempty"`
}

func GetAllURLs() ([]Response, error) {

	rdb := database.CreateClient(0)
	defer rdb.Close()

	keys, err := rdb.Keys(database.Ctx, "*").Result()

	if err != nil {
		return nil, err
	}
	var urls []Response

	for _, key := range keys {
		url, err := rdb.Get(database.Ctx, key).Result()
		if err != nil {
			return nil, err
		}

		// ttl, err := rdb.TTL(database.Ctx, key).Result()

		if err != nil {
			return nil, err
		}

		urls = append(urls, Response{Slug: key, Url: url, Expiry: ""})
	}

	return urls, nil

}
