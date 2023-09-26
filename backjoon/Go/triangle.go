package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, c int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fscanf(r, "%d\n%d\n%d", &a, &b, &c)

	if a+b+c != 180 {
		fmt.Fprintln(w, "Error")
	} else {
		if a == b && b == c {
			fmt.Fprintln(w, "Equilateral")
		} else if a == b || b == c || c == a {
			fmt.Fprintln(w, "Isosceles")
		} else {
			fmt.Fprintln(w, "Scalene")
		}
	}
}