// Simple Go program calculator

package main

import "fmt"

func calc(a, b int) (add_op, sub_op, multi_op, div_op int) {
	add_op = a + b
	sub_op = a - b
	multi_op = a * b
	div_op = a / b
	return
}
func main() {

	var num1, num2 int
	fmt.Scanf("%d %d", &num1, &num2)

	add, sub, multi, div := calc(num1, num2)
	fmt.Printf("The Addition of %d and %d is %d\nThe Subtraction of %d and %d is %d\n", num1, num2, add, num1, num2, sub)
	fmt.Printf("The Division of %d and %d is %d\nThe Multiplication of %d and %d is %d\n", num1, num2, div, num1, num2, multi)
}
