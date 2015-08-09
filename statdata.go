// statdata.go (c) 2015 David Rook - all rights reserved

// calculate count/sum/ave/min/max/sum/sumsq/stdev of data
package statdata

import (
	"fmt"
	"math"
)

type StatDat struct {
	Name   string
	count  int
	sum    float64
	sumsq  float64
	minval float64
	maxval float64
}

func (c *StatDat) Ave() float64 {
	if c.count == 0 {
		return 0
	} // NaN better?
	return c.sum / float64(c.count)
}

func (c *StatDat) Sum() float64 {
	return c.sum
}

func (c *StatDat) Count() int {
	return c.count
}

// StdDev math variety - stat variety uses n-1 divisor if n < 30 if I recall right
// math variety is what a calculator would give you
func (c *StatDat) StdDev() float64 {
	if c.count <= 1 {
		return math.NaN()
	}
	n := float64(c.count)
	variance := (c.sumsq - ((c.sum * c.sum) / n)) / (n)
	return math.Sqrt(variance)
}

// Stat  a datum to the counter
func (c *StatDat) Stat(data float64) {
	//fmt.Printf("New data: %g\n",data)
	if c.count == 0 {
		c.minval = data
		c.maxval = data
	} else {
		if data > c.maxval {
			c.maxval = data
		}
		if data < c.minval {
			c.minval = data
		}
	}
	c.count++
	c.sum += data
	c.sumsq += float64(data) * float64(data)
	//	fmt.Printf("%s : count(%d) sum(%.1f), ave(%.1f) min(%.1f) max(%.1f) sum(%.1f) sumsq(%.1f) stdev(%.5f)\n",
	//		c.Name, c.count, c.sum, float64(c.sum)/float64(c.count), c.minval, c.maxval, c.sum, c.sumsq, c.StdDev())
}

func AbsF64(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func (c *StatDat) Reset() {
	c.count = 0
	c.sum = 0
}

// prints basic data
func (c *StatDat) Dump() {
	if c.count == 0 {
		fmt.Printf("%s has no data\n", c.Name)
		return
	}
	fmt.Printf("%s : count(%d) sum(%.1f), ave(%.1f) min(%.1f) max(%.1f) sum(%.1f) sumsq(%.1f) stdev(%.5f)\n",
		c.Name, c.count, c.sum, float64(c.sum)/float64(c.count), c.minval, c.maxval, c.sum, c.sumsq, c.StdDev())
}
