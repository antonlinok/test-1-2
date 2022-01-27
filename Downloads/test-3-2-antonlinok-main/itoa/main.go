package main

import (
	"bufio"
	"fmt"
	"os"
)

func itoa(n int) string {
	
}
return n
	

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter n: ")
	scanner.Scan()
	n := scanner.Text()
	fmt.Println(itoa(n))
}
