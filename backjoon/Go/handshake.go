package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscanln(reader, &n)
	var k = make([]int, n+1)
	k[1] = 1
	k[2] = 2
	for i := 3; i < n+1; i++ {
		k[i] = (k[i-1] + k[i-2]) % 10
	}
	fmt.Fprintln(writer, k[n])
}