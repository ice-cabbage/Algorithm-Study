package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, count, d int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	
	defer w.Flush()
	fmt.Fscanln(r, &n, &m)
	var factor []int = make([]int, n)

	for i := 1; i <= n; i++ {
		factor[i-1] = n % i
		if n%i == 0 {
			count++
		}
	}

	if count >= m {
		for i := 0; i < n; i++ {
			if factor[i] == 0 {
				d++
				if d == m {
					fmt.Fprintln(w, i+1)
				}
			}
		}
	} else {
		fmt.Fprintln(w, 0)
	}
}