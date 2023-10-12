package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, c int
	fmt.Fscanln(reader, &n, &c)

	var x = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &x[i])
	}

	sort.Ints(x)
	fmt.Fprintln(writer, getRouterDistacne(x, c))
}

func getRouterDistacne(x []int, c int) (dist int) {
	var start = 1                
	var end = x[len(x)-1] - x[0] 

	for start <= end {
		mid := (start + end) / 2
		prev := x[0]
		count := 1

		for i := 1; i < len(x); i++ {
			temp := x[i] - prev
			if mid <= temp {
				count++
				prev = x[i]
			}
		}

		if count >= c {
			dist = mid
			start = mid + 1 
		} else {
			end = mid - 1 
		}
	}
	return
}