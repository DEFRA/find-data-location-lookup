# Location Lookup

This repository contains a simple location lookup API that given a partial location name will return a bounding box for that location.  This should work for towns, villages, cities and 'settlements' that are available in a trimmed down version of the [OpenNames CSV file](https://www.ordnancesurvey.co.uk/opendatadownload/products.html#OPNAME).

## Using the repo

### Setup

1. Copy this repo into your gopath .. `$GOPATH/src/github.com/rossjones/location-lookup`
2. Install [dep](https://github.com/golang/dep)
3. Run `dep ensure`

### The makefile

Running the app

```
export LOCATION_MERGE_FILE=/path/to/csv
make run
```

Building the app

```
export LOCATION_MERGE_FILE=/path/to/csv
make build
```

A sample merge_file (for LOCATION_MERGE_FILE) is available in this repository.