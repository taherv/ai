package client

import "fmt"

func AddNumbers(a int, b int) int {
	return a + b
}

func main() {
	result := AddNumbers(5, 7)
	fmt.Println(result)
}
