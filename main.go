package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
)

const keyServerAddr = "serverAddr"

func getRedirect(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pas_path := strings.Split(r.URL.Path, "/")[1]

	if pas_path == "" {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fs := http.FileServer(http.Dir("static"))
		fs.ServeHTTP(w, r)
		return
	}

	fmt.Printf("%s: got / request. %s \n",
		ctx.Value(keyServerAddr), pas_path,
	)
}
func getHello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /hello request\n", ctx.Value(keyServerAddr))
	io.WriteString(w, "hello route!\n")
}

func main() {
	mux := http.NewServeMux()
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
