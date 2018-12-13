package main

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func csvFileToRecords(csvFile string) []LocationRecord {
	var records []LocationRecord

	in, _ := ioutil.ReadFile(csvFile)
	r := csv.NewReader(strings.NewReader(string(in)))

	for {
		row, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// GlobalType is in cell index 6, no point parsing if we
		// intend to ignore it
		if row[6] != "populatedPlace" {
			continue
		}

		records = append(records, NewLocationRecord(row))
	}

	return records
}
