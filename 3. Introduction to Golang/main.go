package main

import (
	"fmt"
	"intro-golang/library"
)

// Easier variable declaration outside function scope (global variable)
var (
	// Pi is a mathematical constant
	Pi = 3.14
	// Language is the programming language
	Language = "Go"
	// Message is a message
	longer_message = `My name is "John Wick" and I'm a hitman,
	I'm a professional
	killer`
)

// Function swap string value
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	//Shorthand for declaring variable (only work in a function) without data types
	nama := "John Wick"
	fmt.Println(nama)

	// Multi variable declaration
	var first, second, third, fourth string = "satu", "dua", "tiga", "empat"
	fmt.Println(first, second, third, fourth) // Output satu dua tiga empat
	seventh, eight, ninth, ten := "tujuh", "delapan", "sembilan", "sepuluh"
	fmt.Println("Print variables : ", seventh, eight, ninth, ten) //Output tujuh delapan sembilan sepuluh

	// Declaration underscore variable
	_ = "Output"
	_ = "Golang is a fusion language"
	var name string
	name, _ = "john", "wick"
	println(name) // Output = john

	// Decimal numerical
	var floatNumber = 12.365
	fmt.Println(floatNumber)
	fmt.Printf("Decimal number %f\n", floatNumber)   // Output = 12.365000
	fmt.Printf("Decimal number %.4f\n", floatNumber) // Output = 12.3650
	fmt.Printf("Decimal number %e\n", floatNumber)   // Output = 1.236500e+01

	//Boolean
	var exist bool = true
	fmt.Printf("exist? %t\n", exist)

	//String
	var message string = `My name is Jeff
	Golang is pretty awesome tho, you can
	do this`
	fmt.Println(message)
	fmt.Println(longer_message) // Output from global variable declaration

	// Constant variable
	const (
		// Language is the programming language
		Language_2 = "Python"
		// Message is a message
		Message_2 = "Hello World"
	)
	// changing value of constant variable will cause error
	// Language_2 = "Go"
	fmt.Println(Language_2, Message_2)

	// If else statement
	var point int = 5
	if point == 10 {
		fmt.Println("lulus dengan nilai sempurna")
	} else if point >= 5 {
		fmt.Println("Lulus")
	} else if point == 4 {
		fmt.Println("Hampir lulus")
	} else {
		fmt.Println("tidak lulus")
	}

	// Switch statement
	var point_2 int = 5
	switch point_2 {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Switch statement with multiple condition
	var point_3 int = 5
	switch point_3 {
	case 8:
		fmt.Println("perfect")
	case 7, 6, 5, 4:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}

	// Looping with for range
	var fruits = [4]string{"apple", "joe", "mama", "dunno"}
	for i, fruit := range fruits {
		fmt.Printf("elemen ke-%d : %s\n", i, fruit)
	}

	// Looping with for loop
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// Looping with for loop, break and continue
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}
		if i > 8 {
			break
		}
		fmt.Println("Angka", i) // Output 2,4,6,8
	}

	// Calling function swap
	a, b := swap("John", "Wick")
	fmt.Println(a, b)

	// different style to named function parameter but still outside of a function to declare it.
	// func add(x, y int) int {
	// 	return x + y
	// }
	// func add(x int, y int) int {
	// 	return x + y
	// }

	// Multiple output values function
	// func addAndSubs(x, y int) (int, int) {
	// 	return x + y, x - y
	// }

	// Struct and embeded struct
	type person struct {
		name string
		age  int
	}
	type student struct {
		matkul string
		person
	}
	var muridSatu = student{}
	muridSatu.name = "John Wick"
	muridSatu.age = 20
	muridSatu.matkul = "Math"
	fmt.Println("Name : ", muridSatu.name)
	fmt.Println("Age : ", muridSatu.age)
	fmt.Println("Matkul : ", muridSatu.matkul)

	// Combine struct and slice
	var allStudents = []person{
		{name: "John Wick", age: 20},
		{name: "John Wick 2", age: 21},
		{name: "John Wick 3", age: 22},
	}
	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age, "and learn", muridSatu.matkul)
	}

	// Datas structure: Map
	var chicken = map[string]int{}
	chicken["januari"] = 50
	chicken["februari"] = 40
	fmt.Println("januari", chicken["januari"])
	fmt.Println("februari", chicken["februari"])
	fmt.Println("maret", chicken["maret"])

	// Data structure: Array
	var names [4]string
	names[0] = "John"
	names[1] = "Wick"
	names[2] = "Lewis"
	names[3] = "Carroll"
	fmt.Println(names[0], names[1], names[2], names[3])

	// Multidimension array
	var multidim [2][2]int
	multidim[0][0] = 1
	multidim[0][1] = 2
	multidim[1][0] = 4
	multidim[1][1] = 5
	fmt.Println(multidim)

	// Slice
	var fruits_3 = []string{"apple", "grape", "banana", "melon"}
	println(len(fruits_3)) // Output 4
	println(cap(fruits_3)) // Output 4
	fruits_3 = append(fruits_3, "papaya")
	println(len(fruits_3)) // Output 5
	println(cap(fruits_3)) // Output 8
	fmt.Println(fruits_3)
	fmt.Println(fruits_3[0])   // Output apple
	fmt.Println(fruits_3[1])   // Output grape
	fmt.Println(fruits_3[4])   // Output papaya
	fmt.Println(fruits_3[0:2]) // Output [apple grape]
	fmt.Println(fruits_3[1:])  // Output [grape banana melon papaya]
	fmt.Println(fruits_3[:3])  // Output [apple grape banana]

	// Pointer
	var numberA int = 4
	var numberB *int = &numberA
	fmt.Println(numberA) // Output 4
	fmt.Println(numberB) // Output memory address

	// Importing package
	library.SayHello()
	library.SayGoodBye()
	// Defer
	defer fmt.Println("Pembelajaran sudah berakhir")
	fmt.Println("Pembelajaran sudah dimulai")
}
