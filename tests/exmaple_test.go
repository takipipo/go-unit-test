package sum_test

import("fmt"; "go-unit-test/sum")

func ExampleInts(){
	result := sum.Ints(1,2,3,4,5)
	fmt.Println("sum of one to five is", result)
	// Output:
	// sum of one to five is 15
}