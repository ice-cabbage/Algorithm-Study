package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var op string
		fmt.Fscan(reader, &op)
		var n int
		fmt.Fscan(reader, &n)
		var arr string
		fmt.Fscan(reader, &arr)
		var nums []string
		nums = strings.Split(strings.Trim(arr, "[]"), ",")

		front := 0
		back := n
		err := false 

		for j := 0; j < len(op); j++ {
			if op[j] == 'R' {
				front, back = back, front
			} else {
				if n == 0 {
					fmt.Fprintln(writer, "error")
					err = true
					break
				} else {
					if front < back {
						front++
					} else {
						front--
					}
					n--
				}
			}
		}
        
		if err == false {
			if n != 0 {
				fmt.Fprint(writer, "[")
				if front < back {
					for j := front; j < back - 1; j++ {
						fmt.Fprint(writer, nums[j], ",")
					}
					fmt.Fprint(writer, nums[back - 1], "]")
				} else {
					for j := front - 1; j > back; j-- {
						fmt.Fprint(writer, nums[j], ",")
					}
					fmt.Fprint(writer, nums[back], "]")
				}
				fmt.Fprintln(writer)
			} else {
				fmt.Fprintln(writer, "[]")
			}
		}
	}
}