package main

import (
	"fmt"
	"strconv"
	"strings"
)

// LocationRecords is a list of strings
type LocationRecords = []LocationRecord

// SearchResults are, erm, search results
type SearchResults = map[string]LocationRecords

// LocationRecord represents a single row in the CSV files
// Many of the CSV fields are ignored as irrelevant
type LocationRecord struct {
	ID              string  `json:"-"`
	Name            string  `json:"name"`
	Normalised      string  `json:"-"`
	GlobalType      string  `json:"-"`
	LocalType       string  `json:"-"`
	GeometryX       string  `json:"-"`
	GeometryY       string  `json:"-"`
	Xmin            float64 // 12
	Ymin            float64 // 13
	Xmax            float64 // 14
	Ymax            float64 // 15
	MostDetailView  string  // 10
	LeastDetailView string  // 11
	District        string  `json:"district"`
}

// NewLocationRecord creates a new Location record from a list of
// strings (read presumably from a CSV file)
func NewLocationRecord(row []string) LocationRecord {

	loc := LocationRecord{
		ID:              row[0],
		Name:            row[2],
		Normalised:      normalise(row[2]),
		GlobalType:      row[6],
		LocalType:       row[7],
		GeometryX:       row[8],
		GeometryY:       row[9],
		MostDetailView:  row[10],
		LeastDetailView: row[11],
		District:        row[24],
	}

	val, err := strconv.ParseFloat(row[12], 64)
	if err == nil {
		loc.Xmin = val
	}

	val, err = strconv.ParseFloat(row[13], 64)
	if err == nil {
		loc.Ymin = val
	}

	val, err = strconv.ParseFloat(row[14], 64)
	if err == nil {
		loc.Xmax = val
	}

	val, err = strconv.ParseFloat(row[15], 64)
	if err == nil {
		loc.Ymax = val
	}

	return loc
}

func (l LocationRecord) String() string {
	return fmt.Sprintf("{Name:%s, %s\nMin: %f,%f\nMax:%f,%f }\n", l.Name, l.District, l.Xmin, l.Ymin, l.Xmax, l.Ymax)
}

// GetTypeName returns the type of the location
// record provided.
func (l LocationRecord) GetTypeName() string {
	if l.GlobalType == "populatedPlace" {
		return "Place"
	}
	return l.LocalType
}

func normalise(name string) string {
	return strings.ToLower(name)
}
