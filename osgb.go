package main

import (
	"log"

	osgb "github.com/fofanov/go-osgb"
)

var trans osgb.CoordinateTransformer

// MBRToLatLon converts the LocationRecord's osgb to WGS84
// latitude and longitude
func MBRToLatLon(loc LocationRecord) LocationRecord {

	if trans == nil {
		trans, _ = osgb.NewOSTN15Transformer()
	}

	// Min
	minNG := osgb.NewOSGB36Coord(loc.Xmin, loc.Ymin, 0.0)
	gpsCoord, err := trans.FromNationalGrid(minNG)
	if err != nil {
		log.Fatal("1", err)
	}
	loc.Xmin = gpsCoord.Lat
	loc.Ymin = gpsCoord.Lon

	// Max
	maxNG := osgb.NewOSGB36Coord(loc.Xmax, loc.Ymax, 0.0)
	gpsCoord, err = trans.FromNationalGrid(maxNG)
	if err != nil {
		log.Fatal("2", err)
	}
	loc.Xmax = gpsCoord.Lat
	loc.Ymax = gpsCoord.Lon
	return loc
}
