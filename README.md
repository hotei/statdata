<center>
# statdata
</center>

## OVERVIEW

Simple program to collect statistics.

### Installation

If you have a working go installation on a Unix-like OS:

> ```go get github.com/hotei/statdata```

That will copy github.com/hotei/statdata to the first entry of your $GOPATH

or if go is not installed yet :

> ```cd DestinationDirectory```

> ```git clone https://github.com/hotei/statdata.git```

### Features
* simple

### Limitations

* ?

### Usage

See package godoc output

### BUGS
* none known

### Change Log
* 2015-08-09 compiled with 1.5rc1
* started when?

 
### Resources

* [go language reference] [1] 
* [go standard library package docs] [2]
* [Source for statdata] [3]

[1]: http://golang.org/ref/spec/ "go reference spec"
[2]: http://golang.org/pkg/ "go package docs"
[3]: http://github.com/hotei/statdata "github.com/hotei/statdata"

Comments can be sent to <hotei1352@gmail.com> or to user "hotei" at github.com.
License is BSD-two-clause seen below.

### License
```
Released with BSD 2-clause license 

Copyright (c) 2015 - David Rook  <hotei1352@gmail.com>

All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met: 

1. Redistributions of source code must retain the above copyright notice, this
   list of conditions and the following disclaimer. 
2. Redistributions in binary form must reproduce the above copyright notice,
   this list of conditions and the following disclaimer in the documentation
   and/or other materials provided with the distribution. 

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
(INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
```

Documentation (c) 2015 David Rook 

DO NOT EDIT BELOW - Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)

# statdata
    import "."

statdata.go (c) 2015 David Rook - all rights reserved

datatype that can calculate count/sum/ave/min/max/sum/sumsq/stdev of data fed to it


	Note that most data is accessed directly, but Ave() and StdDev() are methods.

This will often save a bit of CPU cycles since typical usage is to feed multiple
data items to the object and then request statistics once.  However, you COULD
recalculate Ave and StdDev after every new datum and put them in new fields
with the appropriate names.

Normally the object will be given a name that will be printed on the Dump()






## func AbsF64
``` go
func AbsF64(a float64) float64
```


## type StatDat
``` go
type StatDat struct {
    Name   string
    Count  int
    Sum    float64
    Sumsq  float64
    Minval float64
    Maxval float64
}
```








### func NewStatDat
``` go
func NewStatDat(title string) *StatDat
```



### func (\*StatDat) Ave
``` go
func (c *StatDat) Ave() float64
```


### func (\*StatDat) Dump
``` go
func (c *StatDat) Dump()
```
prints basic data



### func (\*StatDat) Reset
``` go
func (c *StatDat) Reset()
```
Reset clears data so object is back to its original state



### func (\*StatDat) Stat
``` go
func (c *StatDat) Stat(data float64)
```
Stat  a datum to the counter



### func (\*StatDat) StdDev
``` go
func (c *StatDat) StdDev() float64
```
StdDev math variety - stat variety uses n-1 divisor if n < 30 if I recall right
math variety is what a calculator would give you









- - -
DO NOT EDIT ABOVE - Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)

github.com/hotei/statdata imports 
```
	errors
	fmt
	io
	math
	os
	reflect
	runtime
	strconv
	sync
	sync/atomic
	syscall
	time
	unicode/utf8
	unsafe
```
```
deadcode results:

```
