package durationutil

import (
	"testing"
	"time"
)

func TestToDuration_ValidInputs(t *testing.T) {
	tests := []struct {
		input    string
		expected time.Duration
	}{
		{"1m", time.Minute},
		{"2h", 2 * time.Hour},
		{"3d", 3 * 24 * time.Hour},
		{"1w", 7 * 24 * time.Hour},
		{"2mo", 2 * 30 * 24 * time.Hour},
		{"1y", 365 * 24 * time.Hour},
		{"1y2mo3w4d5h6m", 1*365*24*time.Hour + 2*30*24*time.Hour + 3*7*24*time.Hour + 4*24*time.Hour + 5*time.Hour + 6*time.Minute},
	}

	for _, tt := range tests {
		got, err := ToDuration(tt.input)
		if err != nil {
			t.Errorf("ToDuration(%q) returned error: %v", tt.input, err)
			continue
		}
		if got != tt.expected {
			t.Errorf("ToDuration(%q) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}

func TestToDuration_InvalidInputs(t *testing.T) {
	tests := []string{
		"5x",   // unsupported unit
		"abc",  // invalid format
		"2.5h", // decimals not allowed
		"",     // empty input
		"1h2",  // number with no unit
		"1hh",  // invalid unit
	}

	for _, input := range tests {
		d, err := ToDuration(input)
		if err == nil {
			t.Errorf("ToDuration(%q) expected error but got nil. Returned: %v", input, d)
		}
	}
}
