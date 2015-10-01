/* {{{ Copyright (c) Paul R. Tagliamonte <paultag@gmail.com>, 2015
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE. }}} */

package haversine

import (
	"math"
)

var earthRadiusMetres float64 = 6371000

type Metres float64

func DegreesToRadians(degrees float64) float64 {
	return degrees * math.Pi / 180
}

type Point struct {
	Lat float64
	Lon float64
}

func (p Point) Delta(point Point) Delta {
	return Delta{
		Lat: p.Lat - point.Lat,
		Lon: p.Lon - point.Lon,
	}
}

func (p Point) toRadians() Point {
	return Point{
		Lat: DegreesToRadians(p.Lat),
		Lon: DegreesToRadians(p.Lon),
	}
}

type Delta struct {
	Lat float64
	Lon float64
}

func Distance(origin, position Point) Metres {
	origin = origin.toRadians()
	position = position.toRadians()

	change := origin.Delta(position)

	a := math.Pow(math.Sin(change.Lat/2), 2) + math.Cos(origin.Lat)*math.Cos(position.Lat)*math.Pow(math.Sin(change.Lon/2), 2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return Metres(earthRadiusMetres * c)
}

// vim: foldmethod=marker
