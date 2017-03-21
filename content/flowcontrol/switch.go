// +build OMIT

package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go запущен на ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}
