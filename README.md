# geolib

## About

Geographical functions

## Installation

```bash
$ go get github.com/hub335/geolib
```

## Example

```go
package main

import (
	"fmt"

	"github.com/hub335/geolib"
)

func main() {
	fmt.Printf("Haversine Distance: %2.2f km\n", geolib.HaversineDistance(
		50.116667, 8.683333, 
		52.516667, 13.3833))
}
```

```bash
$ go get github.com/hub335/geolib
$ go run example.go
Haversine Distance: 422.77 km

```

## Thanks

This repo is a fork of the awesome original work contributed by [@alouche](github.com/alouche). 

For more information, see the original repository at:
		https://github.com/alouche/go-geolib
