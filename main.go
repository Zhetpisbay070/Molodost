package main

import "fmt"

func fibonacci(n int) int {
	f1, f2 := 0, 1
	for i := 2; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}

func canHold(w1, l1, w2, l2 int) bool {
	rectangle1 := w1 * l1
	rectangle2 := w2 * l2
	res := rectangle1 > rectangle2
	fmt.Println(res)
	return res
}

func quessNumber(n int) string {
	var res string
	fmt.Println(res)
	switch n {
	case 0:
		fmt.Println("zero")
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("six")
	case 7:
		fmt.Println("seven")
	case 8:
		fmt.Println("eight")
	case 9:
		fmt.Println("nine")
	case 10:
		fmt.Println("ten")
	default:
		fmt.Println("lohi")
	}
	return res
}

func pow(n int, p int) int {
	var res int
	res = 1
	for i := 0; i < p; i++ {
		res = res * n
	}
	return res
}

func main() {
	fmt.Println(fibonacci(10))

	if canHold(10, 20, 9, 19) {
	}

	fmt.Println(quessNumber(0))

	fmt.Println(pow(2, 2))

}
