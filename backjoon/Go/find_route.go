package main

import (
	"bufio"
	"fmt"
	"os"
)

var n int
var bufin *bufio.Reader = bufio.NewReader(os.Stdin)
var bufout *bufio.Writer = bufio.NewWriter(os.Stdout)
var grid [101][101]int
var visited [101][101]int

func bfs(start int) {
	queue := make([]int, 0)
	queue = append(queue, start)

	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 0; i < n; i++ {
			if grid[cur][i] == 1 && visited[start][i] == 0 {
				visited[start][i] = 1
				queue = append(queue, i)
			}
		}
	}
}
func main() {
	defer bufout.Flush()
	fmt.Fscanf(bufin, "%d", &n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fscan(bufin, &grid[i][j])
		}
	}
	for i := 0; i < n; i++ {
		bfs(i)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprintf(bufout, "%d ", visited[i][j])
		}
		fmt.Fprint(bufout, "\n")
	}
}