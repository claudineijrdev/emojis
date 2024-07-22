package main

import (
	"context"
	"slices"
	"testing"

	"github.com/ServiceWeaver/weaver/weavertest"
)

func TestSearcher(t *testing.T) {
	runner := weavertest.Local
	queries := []string{"pig","PIG", "black cat", "goo bar baz"}
	exResponses := [][]string{{"🐖" , "🐗", "🐷", "🐽"},{"🐖" , "🐗", "🐷", "🐽"},{"🐈\u200d⬛"},{}}
	runner.Test(t, func(t *testing.T, searcher Searcher){
		ctx := context.Background()
		for idx, query := range queries {
			emoji, err := searcher.Searcher(ctx,query)
			if err != nil {
				t.Fatal(err)
			}
			if slices.Compare(emoji, exResponses[idx]) != 0 {
				t.Fatalf("emojis %v does not match with expected %v",emoji, exResponses[idx])
			}
		}
	})
}