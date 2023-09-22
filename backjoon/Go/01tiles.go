package main

import (
	"bufio"
	"fmt"
	"os"
)

const len = 1000001
const mod = 15746

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	arr := [len]int{0, 1, 2}
	fill_arr(&arr)
	fmt.Fscan(reader, &n)
	fmt.Fprint(writer, arr[n])
}

func fill_arr(arr *[len]int) {
	for i := 3; i < len; i++ {
		arr[i] = (arr[i - 1] % mod + arr[i - 2] % mod) % mod
	}
}