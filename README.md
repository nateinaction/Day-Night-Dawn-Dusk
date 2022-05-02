# Day Night Dawn Dusk

Given a CSV containing a date, time, lat, and lon, this script will output a human readable time of day (day, night, dawn, dusk). This script considers dawn to be from -6° (civil twilight) to 0° and considers dusk to be from 0° to -6°. The script computes twilight based on location and timezone but not elevation, calculations are made assuming sea level. 

This script was used for a school project to better understand what time of day camera traps captured photos of coyotes and gray foxes.

## Usage
```
go run main.go <path to csv>
```

### Example
```
❯ go run main.go test.csv
night
dawn
day
dusk
night
```