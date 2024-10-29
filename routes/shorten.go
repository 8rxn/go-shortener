package routes

import (
	"fmt"

	"github.com/8rxn/go-shortener/database"
)

func SetShortenedURL(url string, slug string, expiry int32) string {
	rdb := database.CreateClient(0)
	defer rdb.Close()

	fmt.Printf("Setting %s to %s\n", slug, url)
	rdb.Set(database.Ctx, slug, url, 0)
	return slug
}
