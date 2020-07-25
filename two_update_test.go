package defer_demo

func ExampleTwoUpdateDefer() {
	TwoUpdateDefer()
	// Output:
	// deferFunc1 Before return
	// deferFunc Before return
	// panic: error out
}
