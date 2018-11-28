package main

import (
	"sort"
	"strings"

	"github.com/derekparker/trie"
)

var t = trie.New()

var locationTypeWeighting = map[string]int{
	"City":                     11,
	"Town":                     10,
	"Village":                  9,
	"Hamlet":                   8,
	"Other Settlement":         7,
	"Suburban Area":            6,
	"Named Road":               5,
	"Numbered Road":            4,
	"Section Of Named Road":    3,
	"Section Of Numbered Road": 2,
	"Postcode":                 1,
}

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
				return sortSearchResults(searchResults)
			}
		}
	}

	return sortSearchResults(searchResults)
}

func sortSearchResults(searchResults map[string]LocationRecords) map[string]LocationRecords {
	for key := range searchResults {
		sort.Slice(searchResults[key][:], func(i, j int) bool {
			return locationTypeWeighting[searchResults[key][i].LocalType] > locationTypeWeighting[searchResults[key][j].LocalType]
		})
	}
	return searchResults
}
