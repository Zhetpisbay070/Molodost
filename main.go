package main

import "fmt"

func fibonacci(n int) int {
	f1, f2 := 0, 1
	for i := 2; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}

//return n-th number from fibonacci sequence
func main() {
	fmt.Println(fibonacci(6))
}

//func canHold(w1, l1, w2, l2 int) bool {
//	rectangle1 := w1 * l1
//	rectangle2 := w2 * l2
//	res := rectangle1 > rectangle2
//	fmt.Println(res)
//	return res
//}
//
//func main() {
//	if canHold(10, 20, 9, 19) {
//		fmt.Println()
//	}
//}

//func quessNumber(n int) string {
//	var res string
//	fmt.Println(res)
//	return res
//
//}
//
//func main() {
//	number := 5
//	switch number {
//	case 1:
//		fmt.Println("one")
//	default:
//		fmt.Println("lohi")
//	}
//}

//func quessNumber(n int) string {
//	var res string
//	switch n {
//	case 1:
//		res = "one"
//	default:
//		fmt.Println("ebeiii")
//
//	}
//	return res
//}
//func main() {
//	fmt.Println(quessNumber(2))

//	fmt.Println(fibonacci(5))
//}
//
////func fibonacci(n int) int {
////	var res int
////	f1 := 0
////	f2 := 1
////	f3 := f1 + f2
////	res = f2 + f3
////
////	for i := 2; i <= 5; i++ {
////		fmt.Println()
////	}
////
////	//return n-th number from fibonacci sequence
////	return res
