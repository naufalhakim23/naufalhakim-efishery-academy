package main

import (
	"fmt"
	"strconv" // for converting integer to string
)

var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
	minutes      string = "32:23"
	// theURL       string = "https://www.google.com"
	// requestHTTP  string = "GET"
)

// Hello returns a greeting for the named person.
func variables() {
	dark := "bar"

	fmt.Println(dark)
	fmt.Println("Hello, my name is", actorName)
	fmt.Println("I played", companion, "in Doctor Who season", season)
	fmt.Println("I was the", doctorNumber, "doctor")
	fmt.Println("I'm in the", minutes, "mark")
	// Conversion method from int to string
	var (
		i int     = 42
		j float32 = 32.5
	)
	fmt.Printf("%v, %T\n", i, i)
	// fmt.Println()
	fmt.Printf("%v, %T\n", j, j)
	// fmt.Println()
	fmt.Printf("%v, %T\n", i, i)
	var k string = strconv.Itoa(i) // <--- string conversion using strconv I to ascii
	// fmt.Println()
	fmt.Printf("%v, %T\n", k, k)
}
func differentDataTypes() {
	var (
		a int    = 32
		b int8   = 64
		c int16  = 128
		d uint32 = 256 // maximum value of uint32 is 4,294,967,295
	)
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
	// fmt.Println(a+b) --> invalid operation because different data types
	fmt.Println(int8(a) + b) // --> valid operation because same data type
}
func operationInt() {
	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010 = 2
	fmt.Println(a | b)  // 1011 = 11
	fmt.Println(a ^ b)  // 1001 = 9
	fmt.Println(a &^ b) // 0100 = 4
	// bit shifting only in integer
	c := 8              // 2^3
	fmt.Println(c << 3) // 2^3 * 2^3 = 2^6 = 64
	fmt.Println(c >> 3) // 2^3 / 2^3 = 2^0 = 1
}
func floatingPoint() {
	n := 3.14
	fmt.Printf("%v, %T\n", n, n)
	n = 13.7e72
	fmt.Printf("%v, %T\n", n, n)
	n = 2.1e14
	fmt.Printf("%v, %T\n", n, n)
	// but if we want to convert float to int
	// we need to use the math package
	fmt.Println(int64(n))

	a := -12.5
	b := -3.75
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}
func complex() {
	var x complex64 = 1 + 2i
	var y complex64 = 2 + 5.2i
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Println(x + y)
	fmt.Println(x - y)
	fmt.Println(x * y)                       // (1+2i)(2+5.2i) = 2+12.4i+4.2i+10.4i^2 = 2+16.6i-10.4 = -8.2+16.6i
	fmt.Println(x / y)                       // (1+2i)/(2+5.2i) = (1+2i)(2-5.2i)/(2+5.2i)(2-5.2i) = 1/29+4/29i-2/29-10/29i = 29/29-20/29i
	fmt.Printf("%v, %T\n", real(x), real(x)) // -> complex64 is a float32
	fmt.Printf("%v, %T\n", imag(x), imag(x)) // -> complex64 is a float32
	// var z complex128 = complex(5, 12) // -> not working
	// fmt.Printf("%v, %T\n", z, z)
}
func strings() {
	var a string = "this is a string"
	var b string = "much a string"
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", a[2], a[2])                 // -> 105
	fmt.Printf("%v, %T\n", string(a[2]), string(a[2])) // -> i
	fmt.Printf("%v, %T\n", a+b, a+b)                   // -> this is a stringmuch a string
	var c = []byte(a)
	fmt.Printf("%v, %T\n", c, c)
}
func rune() {
	// rune is an alias for int32
	r := 'a'
	fmt.Printf("%v, %T\n", r, r)
}
func main() {
	variables()
	differentDataTypes()
	operationInt()
	floatingPoint()
	complex()
	strings()
	rune()
}
