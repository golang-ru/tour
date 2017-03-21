// +build OMIT

package main

import "fmt"

func main() {
	defer fmt.Println("мир")

	fmt.Println("привет")
}
