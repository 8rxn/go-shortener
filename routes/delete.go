package routes

import (
	"github.com/8rxn/go-shortener/database"
)

func DeleteSlug(slug string) (bool, error) {

	rdb := database.CreateClient(0)
	defer rdb.Close()

	err := rdb.Del(database.Ctx, slug).Err()
	if err != nil {
		return false, err
	}

	return true, nil

}
