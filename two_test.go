package defer_demo

func ExampleTwoDefer() {
	TwoDefer()
	// Output:
	// panic: error out
	// deferFunc1 Before return
	// deferFunc Before return
}
