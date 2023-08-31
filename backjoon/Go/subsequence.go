package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	count int
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, s int
	fmt.Fscanln(reader, &n, &s)
	var seq = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &seq[i])
	}
	computeSubsequenceCount(0, 0, s, seq)
	fmt.Fprintln(writer, count)
}

func computeSubsequenceCount(index, sum, aim int, seq []int) {
	if index > len(seq)-1 {
		return
	}
	sum += seq[index]
	if sum == aim {
		count++
	}
	computeSubsequenceCount(index+1, sum, aim, seq)
	computeSubsequenceCount(index+1, sum-seq[index], aim, seq)
}
