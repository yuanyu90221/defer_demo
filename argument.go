package defer_demo

import "fmt"

func Argument() {
	i := 10
	defer fmt.Println(i)
	i++
	fmt.Println(i)
}
