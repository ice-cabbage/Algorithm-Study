package main

import (
    "bufio"
    "os"
    "strconv"
)

func main() {
    sc := bufio.NewScanner(os.Stdin)
    wr := bufio.NewWriter(os.Stdout)
    defer wr.Flush()

    sc.Scan()
    N, _ := strconv.Atoi(sc.Text())

    que := make([]int, N)
    for i := 1; i < N+1; i++ {
        que[i-1] = i
    }

    for len(que) > 1 {
        que = append(que[2:len(que)], que[1])
    }
    wr.WriteString(strconv.Itoa(que[0]))
}