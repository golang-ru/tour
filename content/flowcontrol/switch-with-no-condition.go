// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Доброе утро!")
	case t.Hour() < 17:
		fmt.Println("Добрый день.")
	default:
		fmt.Println("Добрый вечер")
	}
}
