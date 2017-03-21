// +build OMIT

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Когда пятница?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Сегодня.")
	case today + 1:
		fmt.Println("Завтра.")
	case today + 2:
		fmt.Println("Послезавтра.")
	default:
		fmt.Println("Очень нескоро.")
	}
}
