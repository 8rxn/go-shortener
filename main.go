package main

import (
	"context"
	"encoding/json"
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

type successResponse struct {
	Success bool   `json:"success"`
	Url     string `json:"url"`
	Slug    string `json:"slug"`
	Expiry  string `json:"expiry,omitempty"`
}

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
	var body_json map[string]string
	body, er := io.ReadAll(r.Body)

	if er != nil {
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}

	if err := json.Unmarshal(body, &body_json); err != nil {
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}

	url, ok := body_json["url"]
	if !ok {
		http.Error(w, "url not found in request", http.StatusBadRequest)
		return
	}

	slug, ok := body_json["slug"]
	if !ok {
		http.Error(w, "slug not found in request", http.StatusBadRequest)
		return
	}
	expiry_present := true
	expiry, ok := body_json["expiry"]
	if !ok {
		expiry = "0"
		expiry_present = false
	}
	expiry_time, err := strconv.Atoi(expiry)
	if err != nil {
		http.Error(w, "expiry time not valid", http.StatusBadRequest)
		return
	}

	set_slug := routes.SetShortenedURL(url, slug, int32(expiry_time))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var response successResponse
	response.Success = true
	response.Url = url
	response.Slug = set_slug
	if expiry_present {
		response.Expiry = expiry
	}
	json.NewEncoder(w).Encode(response)
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
