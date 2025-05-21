package durationutil

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ToDuration parses a human-readable duration string (e.g., `2y5mo3d`)
// and returns the corresponding time.Duration.
// Valid time units are: `m` (minute), `h (hour), `d` (day), `w` (week), `mo` (month), and `y` (year)
// Returns an error if the format is invalid or contains unsupported units.
func ToDuration(input string) (time.Duration, error) {
	var duration time.Duration
	r := regexp.MustCompile(`(\d+?)([a-zA-Z]+)`)
	matches := r.FindAllStringSubmatch(input, -1)
	for _, segment := range matches {
		segmentDuration, err := parseDurationSegment(segment[1], segment[2])
		if err != nil {
			return 0, err
		}
		duration =+ segmentDuration
	}

	return duration, nil
}

func parseDurationSegment(unitValue, unit string) (time.Duration, error) {
	var unitDuration time.Duration
	unitToInt, err := strconv.Atoi(unitValue)
	if err != nil {
		return 0, err
	}

	unitDuration = time.Duration(unitToInt)

	switch strings.ToLower(unit) {
	case "m":
		unitDuration = unitDuration * time.Minute
	case "h":
		unitDuration = unitDuration * time.Hour
	case "d":
		unitDuration = unitDuration * time.Hour * 24
	case "w":
		unitDuration = unitDuration * time.Hour * 24 * 7
	case "mo":
		unitDuration = unitDuration * time.Hour * 24 * 30
	case "y":
		unitDuration = unitDuration * time.Hour * 24 * 365
	default:
		return 0, errors.New("couldn't determine duration type of " + unitValue + unit )
	}

	return unitDuration, nil
}