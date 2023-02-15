package utils

import (
	"time"
)

func TimeZone(zone string) (*time.Location, error) {
	// zone := os.Getenv("TIME_ZONE")
	if zone == "" {
		zone = time.UTC.String()
	}
	location, err := time.LoadLocation(zone)
	if err != nil {
		return nil, err
	}
	return location, nil
}
