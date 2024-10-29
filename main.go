package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/8rxn/go-shortener/routes"
)

const keyServerAddr = "serverAddr"

func getRedirect(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	pas_path := strings.Split(r.URL.Path, "/")[1]

	if pas_path == "" {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fs := http.FileServer(http.Dir("static"))
		fs.ServeHTTP(w, r)
		return
	}

	url := routes.GetURL(pas_path)
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
func setUrl(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	url := r.URL.Query().Get("url")
	slug := r.URL.Query().Get("slug")
	expiry := r.URL.Query().Get("expiry")

	if expiry == "" {
		expiry = "0"
	}

	expiry_time, err := strconv.Atoi(expiry)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if url == "" || slug == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	set_slug := routes.SetShortenedURL(url, slug, int32(expiry_time))
	io.WriteString(w, "hello route!\n url: "+url+"\n slug: "+set_slug)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/set", setUrl)
	mux.HandleFunc("/*", getRedirect)

	ctx := context.Background()
	server := &http.Server{
		Addr:    ":5000",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx = context.WithValue(ctx, keyServerAddr, l.Addr().String())
			return ctx
		},
	}

	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}

}
