package main

import (
	"bufio"
	"fmt"
	"os"
)

func equation(a, b float64) {
	if a , b > 0 {
	fmt.Println(-b / a)
	}
	if a, b == 0 {
		fmt.Println(nil)
	}
}
 return

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a: ")
	scanner.Scan()
	a := scanner.Text()
	fmt.Print("Enter b: ")
	scanner.Scan()
	b := scanner.Text()
	fmt.Println(equation(a, b))

}
