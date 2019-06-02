// dem so routine dong thoi voi ham main
package main

import (
	"fmt"
	"runtime"
	"time"
)

func goroutine() {
	fmt.Println("goroutine")
}

func main() {
	go goroutine()
	numberroutine := runtime.NumGoroutine()
	fmt.Println(numberroutine)
	time.Sleep(time.Second)
}
