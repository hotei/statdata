// statdata.go (c) 2015 David Rook - all rights reserved
//
// datatype that can calculate count/sum/ave/min/max/sum/sumsq/stdev of data fed to it
//
//	Note that most data is accessed directly, but Ave() and StdDev() are methods.
//
// This will often save a bit of CPU cycles since typical usage is to feed multiple
// data items to the object and then request statistics once.  However, you COULD
// recalculate Ave and StdDev after every new datum and put them in new fields
// with the appropriate names.
//
// Normally the object will be given a name that will be printed on the Dump()
//
package statdata

import (
	"fmt"
	"math"
)

type StatDat struct {
	Name   string
	Count  int
	Sum    float64
	Sumsq  float64
	Minval float64
	Maxval float64
}

func NewStatDat(title string) *StatDat {
	var statter StatDat
	statter.Name = title
	return &statter
}

func (c *StatDat) Ave() float64 {
	if c.Count == 0 {
		return 0
	} // NaN better?
	return c.Sum / float64(c.Count)
}

// StdDev math variety - stat variety uses n-1 divisor if n < 30 if I recall right
// math variety is what a calculator would give you
func (c *StatDat) StdDev() float64 {
	if c.Count < 1 {
		return math.NaN()
	}
	n := float64(c.Count)
	variance := (c.Sumsq - ((c.Sum * c.Sum) / n)) / (n)
	return math.Sqrt(variance)
}

// Stat  a datum to the counter
func (c *StatDat) Stat(data float64) {
	//fmt.Printf("New data: %g\n",data)
	if c.Count == 0 {
		c.Minval = data
		c.Maxval = data
	} else {
		if data > c.Maxval {
			c.Maxval = data
		}
		if data < c.Minval {
			c.Minval = data
		}
	}
	c.Count++
	c.Sum += data
	c.Sumsq += float64(data) * float64(data)
}

func AbsF64(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

// Reset clears data so object is back to its original state
func (c *StatDat) Reset() {
	c.Count = 0
	c.Sum = 0.0
	c.Sumsq = 0.0
	c.Minval = 0.0 // normally a problem but will be changed by first incoming data
	c.Maxval = 0.0
}

// prints basic data
func (c *StatDat) Dump() {
	if c.Count == 0 {
		fmt.Printf("%s has no data\n", c.Name)
		return
	}
	fmt.Printf("%s : count(%d) sum(%.1f), ave(%.1f) min(%.1f) max(%.1f) sumsq(%.1f) stdev(%.5f)\n",
		c.Name, c.Count, c.Sum, float64(c.Sum)/float64(c.Count), c.Minval, c.Maxval, c.Sumsq, c.StdDev())
}
