package durationutil

import (
	"strconv"
	"strings"
	"time"
)

// HumanizeDuration converts a time.Duration into a human-readable string,
// such as `2 years 5 months 3 days`.
// Any seconds are not included and will be rounded down.
func HumanizeDuration(d time.Duration) string {
	if d == 0 {
		return "0 minutes"
	}

	var parts []string

	for d > 0 {
		switch {
		case d >= time.Hour*24*365:
			count := int(d / (time.Hour * 24 * 365))
			d -= time.Duration(count) * time.Hour * 24 * 365
			unit := "year"
			if count > 1 {
				unit += "s"
			}
			parts = append(parts, strconv.Itoa(count)+" "+unit)
		case d >= time.Hour*24*30:
			count := int(d / (time.Hour * 24 * 30))
			d -= time.Duration(count) * time.Hour * 24 * 30
			unit := "month"
			if count > 1 {
				unit += "s"
			}
			parts = append(parts, strconv.Itoa(count)+" "+unit)
		case d >= time.Hour*24*7:
			count := int(d / (time.Hour * 24 * 7))
			d -= time.Duration(count) * time.Hour * 24 * 7
			unit := "week"
			if count > 1 {
				unit += "s"
			}
			parts = append(parts, strconv.Itoa(count)+" "+unit)
		case d >= time.Hour*24:
			count := int(d / (time.Hour * 24))
			d -= time.Duration(count) * time.Hour * 24
			unit := "day"
			if count > 1 {
				unit += "s"
			}
			parts = append(parts, strconv.Itoa(count)+" "+unit)
		case d >= time.Hour:
			count := int(d / time.Hour)
			d -= time.Duration(count) * time.Hour
			unit := "hour"
			if count > 1 {
				unit += "s"
			}
			parts = append(parts, strconv.Itoa(count)+" "+unit)
		case d >= time.Minute:
			count := int(d / time.Minute)
			d -= time.Duration(count) * time.Minute
			unit := "minute"
			if count > 1 {
				unit += "s"
			}
			parts = append(parts, strconv.Itoa(count)+" "+unit)
		default:
			d = 0
		}
	}

	return strings.Join(parts, " ")
}
