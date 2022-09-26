package main

import (
	libraryFunction "advance-golang/library" // --> import "advance-golang/library" with name libraryFunction
	"fmt"
	"math"
	"sync"
)

// Method
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// or

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	// Pointer --> Pointer is a variable that stores the memory address of another variable
	a := 200
	b := &a
	fmt.Println(a) // Output 200
	fmt.Println(b) // Output memory address
	*b++
	fmt.Println(a) // Output 201

	// Reference Variable --> Reference variable is a variable that stores the memory address of another variable

	// Go Allocation --> Go has automatic allocation and garbage collection

	// Escape Analysis --> Escape analysis is a process that determines whether a variable is allocated on the stack or heap <-- ga ada di materi

	// Go Garbage Collection --> Go has automatic garbage collection <-- ga ada di materi

	// Method --> Method is a function that is associated with a type
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // Output 5
	// or
	fmt.Println(Abs(v)) // Output 5
	// or with type non-struct
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs()) // Output 1.4142135623730951

	// Go Global and Private Variable --> Global variable is a variable that can be accessed from anywhere in the program. Private variable is a variable that can only be accessed from the same package
	println(libraryFunction.ExtractionData(("John Wick")))
	println(libraryFunction.Substraction(10, 5))

	// Go Interface --> Interface is a collection of method signatures that can be implemented by other types.
	l := lingkaran{14}
	s := segitiga{8, 10}
	p := persegi{10}
	fmt.Println("Luas dan Keliling Lingkaran", l.luas(), l.keliling())
	hitungLuasKeliling(l)
	fmt.Println("Luas dan Keliling Segitiga", s.luas(), s.keliling())
	hitungLuasKeliling(s)
	fmt.Println("Luas dan Keliling Persegi", p.luas(), p.keliling())
	hitungLuasKeliling(p)

	// Goroutine --> Go routine is a lightweight thread managed by the Go runtime
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("Hello World")
		}
	}()
	var input string
	fmt.Scanln(&input)

	// Goroutine with WaitGroup --> WaitGroup is a synchronization mechanism that allows you to wait for a collection of goroutines to finish.
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("Hi")
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		wg.Done()
	}()
	wg.Wait()
}

// Go Interface --> Interface is a collection of method signatures that can be implemented by other types.
type lingkaran struct {
	diameter float64
}
type segitiga struct {
	alas   float64
	tinggi float64
}
type persegi struct {
	sisi float64
}
type hitung interface {
	luas() float64
	keliling() float64
}

func (l lingkaran) luas() float64 {
	return 3.14 * l.diameter * l.diameter
}
func (l lingkaran) keliling() float64 {
	return 2 * 3.14 * l.diameter
}
func (s segitiga) luas() float64 {
	return 0.5 * s.alas * s.tinggi
}
func (s segitiga) keliling() float64 {
	return 3 * s.alas
}
func (p persegi) luas() float64 {
	return p.sisi * p.sisi
}
func (p persegi) keliling() float64 {
	return 4 * p.sisi
}
func hitungLuasKeliling(h hitung) {
	fmt.Println("Luas :", h.luas())
	fmt.Println("Keliling :", h.keliling())
}
