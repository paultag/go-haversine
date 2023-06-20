> This library is deprecated in favor of [pault.ag/go/geo](https://github.com/paultag/go-geo), which contains a strict superset of the functionality here.

haversine
=========

This module uses the commonly used haversine formula for calculating small
distances between two points.

Example
-------

```go
package main

import (
	"fmt"
	"pault.ag/go/haversine"
)

func main() {
	whiteHouse := haversine.Point{Lat: 38.89768, Lon: -77.03653}
	eighteenAndF := haversine.Point{Lat: 38.89736, Lon: -77.04173}

	fmt.Printf("%f\n", float64(whiteHouse.MetresTo(eighteenAndF)))
	/* 451.41 Metres (float64) */
}
```
