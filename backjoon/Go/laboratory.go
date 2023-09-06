package main

import "fmt"

type coord struct {
	r int
	c int
}
type node struct {
	node int
	key  int
}

var res = -1

func isEqual(a coord, b coord) bool {
	if a.r == b.r && a.c == b.c {
		return true
	}
	return false
}

func infectVirus(vmap [][]int, r int, c int) {
	if vmap[r][c] == 0 {
		vmap[r][c] = 2
		if r+1 < len(vmap) {
			infectVirus(vmap, r+1, c)
		}
		if c+1 < len(vmap[r]) {
			infectVirus(vmap, r, c+1)
		}
		if r-1 >= 0 {
			infectVirus(vmap, r-1, c)
		}
		if c-1 >= 0 {
			infectVirus(vmap, r, c-1)
		}
	}
}
func search(vmap [][]int, d int, r int, c int) {
	if d == 3 {
		var virus []coord
		count := 0
		for i := range vmap {
			for j := range vmap[i] {
				if vmap[i][j] == 2 {
					var tmp coord
					tmp.r = i
					tmp.c = j
					vmap[i][j] = 0
					virus = append(virus, tmp)
				}
			}
		}
		for _, v := range virus {
			infectVirus(vmap, v.r, v.c)
		}
		for i := range vmap {
			for j := range vmap[i] {
				if vmap[i][j] == 0 {
					count++
				}
			}
		}
		if res < count {
			res = count
		}
		return
	}
	for i := r; i < len(vmap); i++ {
		var j int
		if i == r {
			j = c
		} else {
			j = 0
		}
		for ; j < len(vmap[i]); j++ {
			newVmap := make([][]int, len(vmap))
			for ii := range newVmap {
				newVmap[ii] = append([]int{}, vmap[ii]...)
			}
			if newVmap[i][j] == 0 {
				newVmap[i][j] = 1
				search(newVmap, d+1, i, j)
			}
		}
	}
}

func main() {
	var n, m int
	var virusMap [][]int
	fmt.Scan(&n)
	fmt.Scan(&m)
	virusMap = make([][]int, n)
	
	for i := 0; i < n; i++ {
		virusMap[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&virusMap[i][j])
		}
	}
	search(virusMap, 0, 0, 0)
	fmt.Printf("%d\n", res)
}