package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var wtr = bufio.NewWriter(os.Stdout)

func main() {

	defer flush()
	a := ni()
	o := 0
	if a <= 999 {
		out(o)
		return
	}
	o += a - 999
	if a <= 999999 {
		out(o)
		return
	}
	o += a - 999999
	if a <= 999999999 {
		out(o)
		return
	}
	o += a - 999999999
	if a <= 999999999999 {
		out(o)
		return
	}
	o += a - 999999999999
	if a <= 999999999999999 {
		out(o)
		return
	}
	o += a - 999999999999999
	if a <= 999999999999999999 {
		out(o)
		return
	}
	o += a - 999999999999999999

	out(o)
}

func init() {
	sc.Split(bufio.ScanWords)
	if len(os.Args) > 1 && os.Args[1] == "i" {
		b, e := ioutil.ReadFile("./input")
		if e != nil {
			panic(e)
		}
		sc = bufio.NewScanner(strings.NewReader(strings.Replace(string(b), " ", "\n", -1)))
	}
}

func ni() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nf() float64 {
	sc.Scan()
	f, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return f
}

func ns() string {
	sc.Scan()
	return sc.Text()
}

func out(v interface{}) {
	_, e := fmt.Fprintln(wtr, v)
	if e != nil {
		panic(e)
	}
}

func flush() {
	e := wtr.Flush()
	if e != nil {
		panic(e)
	}
}
