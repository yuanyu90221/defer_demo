package defer_demo

func ExampleTwo() {
	TwoDefer()
	// Output:
	// panic: error out
	// deferFunc1 Before return
	// deferFunc Before return
}
