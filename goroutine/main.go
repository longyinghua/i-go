package main

import (
	"fmt"
	"runtime"
)

func main() {
	n := runtime.GOMAXPROCS(4) //指定以单核运算
	fmt.Println("n = ", n)

	for {
		go fmt.Print(1)

		fmt.Print(0)
	}

}
