package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/nateinaction/dusk/pkg/dusk"
	tzm "github.com/zsefvlol/timezonemapper"
)

func main() {
	filename := os.Args[1]
	locations, err := ImportLocations(filename)
	if err != nil {
		panic(err)
	}

	for _, location := range locations {
		// fmt.Printf("%v\n", location)
		date := time.Date(location.Datetime.Year(), location.Datetime.Month(), location.Datetime.Day(), 0, 0, 0, 0, location.Datetime.Location())
		civilTwilight := dusk.GetSunriseSunsetTimesInUTC(date, -6, location.Lon, location.Lat, 0)
		horizon := dusk.GetSunriseSunsetTimesInUTC(date, 0, location.Lon, location.Lat, 0)
		// fmt.Printf("%v\n", civilTwilight.Rise.In(location.Location))
		// fmt.Printf("%v\n", horizon.Rise.In(location.Location))
		// fmt.Printf("%v\n", horizon.Set.In(location.Location))
		// fmt.Printf("%v\n\n", civilTwilight.Set.In(location.Location))

		if location.Datetime.Before(civilTwilight.Rise.In(location.Location)) {
			fmt.Printf("%v\n", "night")
		}
		if location.Datetime.After(civilTwilight.Rise.In(location.Location)) && location.Datetime.Before(horizon.Rise.In(location.Location)) {
			fmt.Printf("%v\n", "dawn")
		}
		if location.Datetime.After(horizon.Rise.In(location.Location)) && location.Datetime.Before(horizon.Set.In(location.Location)) {
			fmt.Printf("%v\n", "day")
		}
		if location.Datetime.After(horizon.Set.In(location.Location)) && location.Datetime.Before(civilTwilight.Set.In(location.Location)) {
			fmt.Printf("%v\n", "dusk")
		}
		if location.Datetime.After(civilTwilight.Set.In(location.Location)) {
			fmt.Printf("%v\n", "night")
		}
	}
}

type Location struct {
	Datetime time.Time
	Lat      float64
	Lon      float64
	Location *time.Location
}

func ImportLocations(filename string) ([]Location, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var location []Location
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		lat, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return nil, err
		}

		lon, err := strconv.ParseFloat(record[2], 64)
		if err != nil {
			return nil, err
		}

		tz := tzm.LatLngToTimezoneString(lat, lon)
		loc, err := time.LoadLocation(tz)
		if err != nil {
			return nil, err
		}

		t, err := time.ParseInLocation("1/2/06 15:04", record[7], loc)
		if err != nil {
			return nil, err
		}

		location = append(location, Location{
			Datetime: t,
			Lat:      lat,
			Lon:      lon,
			Location: loc,
		})
	}
	return location, nil
}
