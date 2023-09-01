package main

import (
	"bufio"
	"fmt"
	"os"
)

func multiply(a, b, c int) int {
	if b == 1 {
		return a % c
	}
    
	if b%2 == 0 {
		ans := multiply(a, b/2, c)
		return (ans * ans) % c
	}
	return (multiply(a, b/2, c) * multiply(a, b/2+1, c)) % c
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c int
	fmt.Fscan(reader, &a, &b, &c)

	res := multiply(a, b, c)
	fmt.Fprint(writer, res)

}