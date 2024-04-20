package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

var st []bool
var dt []int

var g []map[int]int
var n int

func main() {
	var m int
	fmt.Scanf("%d%d", &n, &m)
	st = make([]bool, n+1)
	dt = make([]int, n+1)
	g = make([]map[int]int, n+1)
	for i := 1; i <= n; i++ {
		g[i] = map[int]int{}
		dt[i] = math.MaxInt64
	}
	dt[1] = 0
	r := bufio.NewReader(os.Stdin)
	for i := 1; i <= m; i++ {
		in := readLine(r)
		x, y, z := in[0], in[1], in[2]
		if _, ok := g[x][y]; ok {
			g[x][y] = min(g[x][y], z)
		} else {
			g[x][y] = z
		}
	}
	res := dijkstra()
	fmt.Println(res)
}
func readline(r *bufio.Reader) []int {
	s, _ := r.ReadString('\n')
	ss := strings.Fields(s)
	res := make([]int, len(ss))
	for i, v := range ss {
		res[i], _ = strconv.Atoi(v)
	}
	return res
}

func readLine(reader *bufio.Reader) []int {
	readString, _ := reader.ReadString('\n')
	ss := strings.Split(readString, " ")
	// ss := strings.Fields(readString)
	var res []int
	for _, s := range ss {
		tmp, _ := strconv.Atoi(strings.TrimSpace(s))
		res = append(res, tmp)
	}
	return res
}

func dijkstra() int {
	var h heaps
	heap.Init(&h)
	heap.Push(&h, Heap{0, 1})
	for h.Len() > 0 {
		top := (heap.Pop(&h)).(Heap)
		if st[top.v] {
			continue
		}
		st[top.v] = true

		for y, z := range g[top.v] {
			if dt[top.v]+z < dt[y] {
				dt[y] = dt[top.v] + z
				heap.Push(&h, Heap{dt[y], y})
			}
		}
	}

	if dt[n] == math.MaxInt64 {
		return -1
	}
	return dt[n]
}

type Heap struct {
	dis int
	v   int
}

type heaps []Heap

func (h heaps) Less(i, j int) bool {
	return h[i].dis < h[j].dis
}

func (h heaps) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h heaps) Len() int {
	return len(h)
}

func (h *heaps) Push(x interface{}) {
	*h = append(*h, x.(Heap))
}

func (h *heaps) Pop() interface{} {
	a := *h
	top := a[len(a)-1]
	*h = a[:len(a)-1]
	return top
}
