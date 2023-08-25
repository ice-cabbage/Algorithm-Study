package main

import (
	"bufio"
	"sort"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, m int
	fmt.Fscanln(reader, &n)
	var cards = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%d ", &cards[i])
	}

	sort.Ints(cards)

	fmt.Fscanln(reader, &m)
	var targets = make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%d ", &targets[i])
		fmt.Fprintf(writer, "%d ", upperBound(cards, targets[i])-lowerBound(cards, targets[i])+1)
	}
	fmt.Fprintln(writer, "")
}

func lowerBound(array []int, target int) int {
	var startIdx, endIdx, midIdx int
	endIdx = len(array) - 1

	for startIdx < endIdx {
		midIdx = (startIdx + endIdx) / 2
		if array[midIdx] >= target {
			endIdx = midIdx
		} else {
			startIdx = midIdx + 1
		}
	}
	return endIdx
}

func upperBound(array []int, target int) int {
	var startIdx, endIdx, midIdx int
	endIdx = len(array) - 1

	for startIdx < endIdx {
		midIdx = (startIdx + endIdx) / 2
		if array[midIdx] > target {
			endIdx = midIdx
		} else {
			startIdx = midIdx + 1
		}
	}

	if array[endIdx] != target {
		endIdx--
	}
	return endIdx
}