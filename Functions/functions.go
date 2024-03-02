package main

import "fmt"

func add(a, b int) {
	fmt.Println("Sum is", a+b)
}
func substract(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("Functions")
	add(5, 6)

	ans := substract(4, 2)
	fmt.Println("Substraction is", ans)

	ans1, ans2 := multiplication(2, 3)
	fmt.Println(ans1, ans2)

	ans3, ans4 := division(12, 24)
	fmt.Println(ans3, ans4)

	advancefunctions()

}

func multiplication(a, b int) (int, int) {
	return a * b, a * b
}

func division(a, b int) (z1, z2 int) {
	z1 = a / b
	z2 = b / a
	return
}
