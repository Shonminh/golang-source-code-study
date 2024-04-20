package main

import (
	"container/heap"
	"fmt"
	"math"
)

// 有一个容器V中存放了海量的三维坐标点 P{double x; double y; double z},点的数量非常巨大。
// 现在已知一个既定的坐标P0，
// P0不在容器V中，请从V中将几何距离P0最近的前n个点删除，
// 其余点保留着V内。
// （使用的语言可以自己选择，容器V可以自己定义，坐标点P可以根据自己的语言声明）

type Vector struct {
	m map[string]*Point
}

func NewVector() *Vector {
	return &Vector{map[string]*Point{}}
}

func (v *Vector) Put(p Point) {
	v.m[p.pointId] = &p
}

func (v *Vector) GetTopN(p Location, n int) []*Point {
	h := &Heaps{p: &p}
	heap.Init(h)
	for _, point := range v.m {
		heap.Push(h, point)
		if h.Len() > n {
			heap.Pop(h)
		}
	}
	var res []*Point
	for h.Len() > 0 {
		res = append(res, (heap.Pop(h)).(*Point))
	}
	return res
}

func (v *Vector) Delete(ps []*Point) {
	for _, p := range ps {
		delete(v.m, p.pointId)
	}
}

func (v *Vector) DeleteTopN(p Location, n int) {
	v.Delete(v.GetTopN(p, n))
}

type Point struct {
	pointId string
	loc     *Location
}

type Location struct {
	x, y, z float64
}

func (l *Location) Distance(p *Location) float64 {
	tmp := math.Pow(l.x-p.x, 2) + math.Pow(l.y-p.y, 2) + math.Pow(l.z-p.z, 2)
	return math.Sqrt(tmp)
}

type Heaps struct {
	p  *Location
	li []*Point
}

func (h Heaps) Len() int {
	return len(h.li)
}

func (h Heaps) Swap(x, y int) {
	h.li[x], h.li[y] = h.li[y], h.li[x]
}

func (h Heaps) Less(x, y int) bool {
	return h.p.Distance(h.li[x].loc) > h.p.Distance(h.li[y].loc)
}

func (h *Heaps) Push(x interface{}) {
	h.li = append(h.li, x.(*Point))
}

func (h *Heaps) Pop() interface{} {
	x := h.li
	top := x[len(x)-1]
	h.li = x[:len(x)-1]
	return top
}

func main() {
	vector := NewVector()
	vector.Put(Point{"北京", &Location{1.0, 1.0, 1.0}})
	vector.Put(Point{"北京2", &Location{1.0, 1.0, 2.0}})
	vector.Put(Point{"北京3", &Location{1.0, 1.0, 3.0}})
	p := Location{0.0, 0.0, 0.0}
	vector.DeleteTopN(p, 3)
	fmt.Println(vector.m)
}
