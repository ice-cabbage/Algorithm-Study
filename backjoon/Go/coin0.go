package main

import "fmt"

func main() {
	var n, k int
	fmt.Scanf("%d %d\n", &n, &k)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d\n", &a[i])
	}

	var cnt int = 0
	for i := n - 1; i >= 0; i-- {
		cnt += k / a[i]
		k %= a[i]
		if k == 0 {
			break
		}
	}
	fmt.Println(cnt)
}
