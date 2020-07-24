package defer_demo

import "fmt"

func deferFun() {
	fmt.Println("deferFunc Before return")
}
func deferFun1() {
	fmt.Println("deferFunc1 Before return")
}

func runPanic() {
	panic(fmt.Errorf("error out"))
}
func TwoDefer() {
	defer deferFun()
	defer deferFun1()

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()
	panic(fmt.Errorf("error out"))
	// runPanic()
	// defer func() {
	// 	runPanic()
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()
}
