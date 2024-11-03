package routes

import (
	"github.com/8rxn/go-shortener/database"
)

func GetAllURLs() (map[string]string, error) {

	rdb := database.CreateClient(0)
	defer rdb.Close()

	keys, err := rdb.Keys(database.Ctx, "*").Result()

	if err != nil {
		return nil, err
	}
	urls := make(map[string]string)

	for i, key := range keys {
		keys[i], err = rdb.Get(database.Ctx, key).Result()
		if err != nil {
			return nil, err
		}
		urls[key] = keys[i]
	}

	return urls, nil

}
