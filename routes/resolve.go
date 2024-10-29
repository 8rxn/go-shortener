package routes

import (
	"github.com/8rxn/go-shortener/database"
)

func GetURL(slug string) string {

	rdb := database.CreateClient(0)
	defer rdb.Close()

	val, err := rdb.Get(database.Ctx, slug).Result()
	if err != nil {
		return ""
	}
	return val
}
