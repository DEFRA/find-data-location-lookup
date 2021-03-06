package main

import (
	"fmt"
	"os"
	"time"
)

// MaxResults is the maximum number of records to return
// in response to a user initiated search
const MaxResults = 20

func main() {
	fmt.Println("Location lookup service loading ...")

	filepath := os.Getenv("LOCATION_MERGE_FILE")
	if filepath == "" {
		fmt.Println("You must specify LOCATION_MERGE_FILE")
		os.Exit(1)
	}

	count := 0
	start := time.Now()
	records := csvFileToRecords(filepath)
	for _, record := range records {
		// Find an existing list in the tree, if not there add a new one
		node, exists := locationTrie.Find(record.Normalised)
		var location []LocationRecord

		if !exists {
			location = []LocationRecord{record}
		} else {
			location = node.Meta().([]LocationRecord)
			location = append(location, record)
		}
		locationTrie.Add(record.Normalised, location)

		count = count + 1
	}
	fmt.Println("..data loaded ..", count, " records in ", time.Since(start))

	runserver()
}
