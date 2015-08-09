// statdat_test.go (c) 2015 David Rook - all rights reserved

package statdata

import (
	"fmt"
	"testing"
)

func Test001(t *testing.T) {
	var ct StatDat = StatDat{Name: "MyStats"}

	ct.Stat(1.0)
	ct.Stat(2.0)
	ct.Stat(3.0)
	ct.Stat(1.0)

	fmt.Printf("Final result follows:\n")
	ct.Dump()
	// compare to stdev value calculated by Sharp EL-533 calculator
	el533 := .829156197
	if AbsF64((ct.StdDev() - el533)) > 0.000001 {
		t.Fatalf("Test001 failed at StdDev()\n")
	}
}
