// +build OMIT

package main

import "fmt"

func main() {
	fmt.Println("подсчет")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("готово")
}
