package defer_demo

import "fmt"

func TwoUpdateDefer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()
	defer fmt.Println("deferFunc Before return")
	defer fmt.Println("deferFunc1 Before return")
	panic(fmt.Errorf("error out"))
}
