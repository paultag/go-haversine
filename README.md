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
	fmt.Printf("%f\n", float64(
		haversine.Point{ /* White House */
			Lat: 38.89768,
			Lon: -77.03653,
		}.MetresTo(
			haversine.Point{ /* 18th & F St, NE */
				Lat: 38.89736,
				Lon: -77.04173,
			},
		),
	)) /* 451.41 Metres (float64) */
}
```
