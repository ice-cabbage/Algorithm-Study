package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	r = bufio.NewReader(os.Stdin)
	w = bufio.NewWriter(os.Stdout)
	n int
	arr []int
)

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func max() int {
	num, idx := 0, 0
	for i := 0; i < len(arr); i++ {
		if num < arr[i] {
			num, idx = arr[i], i
		}
	}
	return idx
}

func main() {
	
	defer w.Flush()
	fmt.Fscan(r, &n)

	for i := 0; i < n; i++ {
		var a int
		fmt.Fscan(r, &a)
		arr = append(arr, a)
	}

	ans := 0
	for i := 0; i < n-1; i++ {
		idx := max()
		if idx == 0 {
			ans += Abs(arr[idx] - arr[idx+1])
			arr = arr[1:]
		} else if idx == len(arr)-1 {
			ans += Abs(arr[idx] - arr[idx-1])
			arr = arr[:len(arr)-1]
		} else {
			var num int
			if arr[idx-1] > arr[idx+1] {
				num = arr[idx-1]
			} else {
				num = arr[idx+1]
			}
			ans += Abs(arr[idx] - num)
			arr = append(arr[:idx], arr[idx+1:]...)
		}
	}
	fmt.Fprintln(w, ans)
}