package library

import "fmt"

// Global function
func SayHello() {
	fmt.Println("hello")
	introduce("John Wick")
}
func SayGoodBye() {
	fmt.Println("good bye")
}

// Private function
func introduce(name string) {
	fmt.Println("My name is", name)

}
