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

	var cnt = make([]int, n+1)
	for i := 1; i < n+1; i++ {
		if i == 1 {
			cnt[i] = 1
		} else if i == 2 {
			cnt[i] = 3
		} else {
			cnt[i] = (cnt[i-1] + 2*cnt[i-2]) % 10007
		}
	}
	fmt.Fprintln(writer, cnt[n])
}
