package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "embed"

	"github.com/ServiceWeaver/weaver"
)

type app struct {
	weaver.Implements[weaver.Main]
	searcher weaver.Ref[Searcher]
	emojis   weaver.Listener
}

//go:embed index.html
var indexHtml string

func run(ctx context.Context, a *app) error {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		if r.URL.Path != "/" {
			http.NotFound(w,r)
			return
		}

		fmt.Fprint(w, indexHtml)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		emojis, err := a.searcher.Get().Searcher(ctx, query)
		if err != nil {
			log.Fatal(err)
		}


		bytes, err := json.Marshal(emojis)
		if err != nil {
			a.Logger(r.Context()).Error("error marshaling search results", "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Println(string(bytes))
		if _, err := fmt.Fprintln(w, string(bytes)); err != nil {
			a.Logger(r.Context()).Error("error writing search results", "err", err)
		}
	})
	return http.Serve(a.emojis,nil)
}
func main() {
	if err := weaver.Run(context.Background(), run); err != nil {
		panic(err)
	}
}
