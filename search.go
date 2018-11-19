package main

import (
	"strings"

	"github.com/derekparker/trie"
)

var t = trie.New()

func getSearchResults(term string) SearchResults {
	searchResults := make(map[string]LocationRecords)

	counter := 0

	res := t.PrefixSearch(strings.ToLower(term))
	for _, item := range res {
		node, _ := t.Find(item)
		meta := node.Meta()
		recordList := meta.([]LocationRecord)

		for _, lr := range recordList {
			t := lr.GetTypeDisplay()
			lr = MBRToLatLon(lr)
			searchResults[t] = append(searchResults[t], lr)
			counter++
			if counter == MaxResults {
				return searchResults
			}
		}
	}

	return searchResults
}
