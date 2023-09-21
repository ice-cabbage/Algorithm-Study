package main

import (
	"bufio"
	"fmt"
	"os"
)

var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)

var s string
var index int 

func main(){
	defer bufout.Flush()
	fmt.Fscan(bufin,&s)
	index = 0 
	if len(s) % 3 == 1 {
		s = s[:index] + "00" + s[index:]
	}else if len(s) % 3 == 2 {
		s = s[:index] + "0" + s[index:]
	}
	for i := 0; i < len(s); i = i + 3 {
		fmt.Fprintf(bufout, "%d",(s[i] - '0') * 4 + (s[i + 1] - '0') * 2 + (s[i + 2] - '0'))
	}
}