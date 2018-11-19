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

		record := NewLocationRecord(row)
		if record.GlobalType != "populatedPlace" {
			continue
		}

		records = append(records, record)
	}

	return records
}
