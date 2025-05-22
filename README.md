# durationutil

A lightweight Go utility package that provides human-readable and compact conversions between `time.Duration` values and human-friendly string representations.

[![MIT license](https://img.shields.io/badge/license-MIT-brightgreen.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/RhykerWells/durationutil)](https://goreportcard.com/report/github.com/RhykerWells/durationutil)
[![Go Reference](https://pkg.go.dev/badge/RhykerWells/durationutil.svg)](https://pkg.go.dev/github.com/RhykerWells/durationutil)
![GitHub Release](https://img.shields.io/github/v/release/RhykerWells/durationutil)


## Features

- Convert `time.Duration` into readable strings like `2 years 5 months 3 days` 
- Parse shorthanded go-like durations such as `2y5mo3d` into valid `time.Duration`

## Usage
To install the package, use `go get` to install it
```bash
go get github.com/RhykerWells/durationutil
```

Here's an example on how to use this package in your project.
```go
import (
	"fmt"
	"time"
	"github.com/RhykerWells/durationutil"
)

func main() {
	// Example 1: HumanizeDuration
	d := 48*time.Hour + 90*time.Minute
	fmt.Println(durationutil.HumanizeDuration(d))
    // Output: "2 days 1 hour 30 minutes"

	// Example 2: ToDuration
	d2, err := durationutil.ToDuration("1y2mo3w4d5h6m")
	if err != nil {
		panic(err)
	}
	fmt.Println(d2) // Output: "10805h6m0s"
}
```