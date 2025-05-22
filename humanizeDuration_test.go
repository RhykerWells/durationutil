package durationutil

import (
	"testing"
	"time"
)

func TestHumanizeDuration(t *testing.T) {
	tests := []struct {
		input    time.Duration
		expected string
	}{
		{0, "0 minutes"},
		{time.Minute, "1 minute"},
		{5 * time.Minute, "5 minutes"},
		{time.Hour, "1 hour"},
		{3 * time.Hour, "3 hours"},
		{24 * time.Hour, "1 day"},
		{48 * time.Hour, "2 days"},
		{7 * 24 * time.Hour, "1 week"},
		{3 * 7 * 24 * time.Hour, "3 weeks"},
		{30 * 24 * time.Hour, "1 month"},
		{60 * 24 * time.Hour, "2 months"},
		{365 * 24 * time.Hour, "1 year"},
		{3 * 365 * 24 * time.Hour, "3 years"},
		{(365*24*time.Hour + 60*24*time.Hour + 3*24*time.Hour + 4*time.Hour + 5*time.Minute), "1 year 2 months 3 days 4 hours 5 minutes"},
		{(365+30+7+1)*24*time.Hour + time.Hour + time.Minute, "1 year 1 month 1 week 1 day 1 hour 1 minute"},
	}

	for _, tt := range tests {
		result := HumanizeDuration(tt.input)
		if result != tt.expected {
			t.Errorf("HumanizeDuration(%v) = %q, want %q", tt.input, result, tt.expected)
		}
	}
}