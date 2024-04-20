package main

import (
	"fmt"
)

// 输入一个以字符串表示的非负整数 sNum，从这个串中移除 k 位数字，使得剩下的数字最小。
// 例如：
// 输入: num = "123421", k = 2； 输出: "1221"
// 输入: num = "10300", k = 1； 输出: "300"
// 输入: num = "1432219", k = 3，输出: "1219"

func findMin2(s string, k int) string {
	var ns = []byte(s)
	var st []byte
	for i := 0; i < len(ns); i++ {
		if k == 0 || len(st) == 0 {
			st = append(st, ns[i])
			continue
		}
		for k > 0 && len(st) > 0 && st[len(st)-1] >= ns[i] {
			st = st[:len(st)-1]
			k--
		}
		st = append(st, ns[i])
	}

	for k > 0 {
		st = st[:len(st)-1]
		k--
	}

	j := 0
	for st[j] == '0' {
		j++
	}
	st = st[j:]
	if len(st) == 0 {
		return "0"
	}
	return string(st)
}

func main() {
	r := findMin("123421", 2)
	fmt.Println(r)
	r = findMin("10300", 1)
	fmt.Println(r)
	r = findMin("1442219", 3)
	fmt.Println(r)
	r = findMin("123", 1)
	fmt.Println(r)
}

var res [][]byte

var path []byte

func dfs(ns []byte, memo []bool, index int, k int) {
	if len(path) == len(ns)-k {
		tmp := make([]byte, len(path))
		copy(tmp, path)
		res = append(res, tmp)
		return
	}

	for i := index; i < len(ns); i++ {
		if memo[i] {
			continue
		}
		path = append(path, ns[i])
		memo[i] = true
		dfs(ns, memo, i+1, k)
		path = path[:len(path)-1]
		memo[i] = false
	}
}

func findMin(s string, k int) string {
	res = nil
	path = nil
	ns := []byte(s)
	memo := make([]bool, len(ns))
	dfs(ns, memo, 0, k)
	var val = string(res[0])
	for _, bs := range res {
		val = min(string(bs), val)
	}
	return val
}

func min(a, b string) string {
	if a > b {
		return b
	}
	return a
}

// // 生男孩就不升了，生女孩  整个国家那女
// 男  女
// 8  8
//
//	4  4
//	   2  2
//	      1  1
//	          1  1
//
// 男 == 女
type T struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			K  int     `json:"k"`
			O  string  `json:"o"`
			C  string  `json:"c"`
			H  string  `json:"h"`
			L  string  `json:"l"`
			V  int     `json:"v"`
			T  float64 `json:"t"`
			R  float64 `json:"r"`
			Lc float64 `json:"lc"`
			Cp string  `json:"cp"`
		} `json:"list"`
	} `json:"data"`
}
