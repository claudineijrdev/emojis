package main

import (
	"context"
	"slices"
	"sort"
	"strings"

	"github.com/ServiceWeaver/weaver"
)

type Searcher interface {
	Searcher(ctx context.Context, query string) ([]string, error)
}

type searcher struct {
	weaver.Implements[Searcher]
	cache weaver.Ref[Cache]
}

func matches(labels, words []string) bool {
	for _, word := range words {
		if !slices.Contains(labels, word) {
			return false
		}
	}
	return true
}

func (s *searcher) Searcher(ctx context.Context, query string) ([]string, error) {
	s.Logger(ctx).Debug("Search", "query", query)

	c, err := s.cache.Get().Get(ctx,query)
	if err != nil {
		s.Logger(ctx).Error("Search", "Error on getting cache data", err)
	}

	if c != nil {
		return c, nil
	}

	s.Logger(ctx).Debug("Getting data from db", "query",query)
	queryArgs := strings.Split(strings.ToLower(query), " ")
	var response []string

	for emoji, labels := range emojis {
		if matches(labels, queryArgs) {
			response = append(response, emoji)
		}
	}

	sort.Strings(response)
	err = s.cache.Get().Put(ctx,query,response)
	if err != nil {
		s.Logger(ctx).Error("Search", "Error on caching data", err)
	}
	return response, nil
}