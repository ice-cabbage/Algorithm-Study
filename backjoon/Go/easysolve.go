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

	var a, b int
	fmt.Fscanln(reader, &a, &b)
	sequence := make([]int, b)
	num := 1
	index := 0
	for index < b {
		j := 0
		for ; j < num && j < b-index; j++ {
			sequence[index+j] = num
		}
		num++
		index += j
	}
	var result int
	for i := a - 1; i < b; i++ {
		result += sequence[i]
	}
	fmt.Fprintln(writer, result)
}