package main

import (
	"bytes"
	"fmt"
	"sort"
)

var (
	o     = "o"
	x     = "x"
	d3map = make(map[string]bool)
	d4map = make(map[string]bool)
)

func d2(n, m int, test func(int, int) bool) {
	var os, xs int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if test(i, j) {
				fmt.Print(o)
				os++
			} else {
				fmt.Print(x)
				xs++
			}
		}
		fmt.Print("\n")
	}
	fmt.Printf("%d os, %d xs\n", os, xs)
}

// n -> rows, m -> blocks, l -> columns
func d3(n, m, l int, test func(int, int, int) bool) {
	var os, xs int
	rows := make([]bytes.Buffer, m)
	// this is going to make rows the outer loop as one would expect
	// for each row, for each block, loop over columns, append that value to the buffer corresponding to the row, between blocks put |
	// rows
	for i := 0; i < n; i++ {
		// blocks
		for j := 0; j < m; j++ {
			// columns
			for k := 0; k < l; k++ {
				if test(i, j, k) {
					fmt.Fprint(&rows[i], o)
					os++
				} else {
					fmt.Fprint(&rows[i], x)
					xs++
				}
			}
			if j < m-1 {
				fmt.Fprint(&rows[i], "|")
			}
		}
		fmt.Println(rows[i].String())
		rows[i].Reset()
	}
	fmt.Printf("%d os, %d xs\n", os, xs)
}

func d4(q, n, m, l int, test func(int, int, int, int) bool) {
	rows := make([]bytes.Buffer, m)
	var os, xs int
	// just like above but repeat q times, separate by line of "-"
	for h := 0; h < q; h++ {
		for i := 0; i < n; i++ {
			// blocks
			for j := 0; j < m; j++ {
				// columns
				for k := 0; k < l; k++ {
					if test(h, i, j, k) {
						fmt.Fprint(&rows[i], o)
						os++
					} else {
						fmt.Fprint(&rows[i], x)
						xs++
					}
				}
				if j < m-1 {
					fmt.Fprint(&rows[i], "|")
				}
			}
			fmt.Println(rows[i].String())
			rows[i].Reset()
		}
		if h < q-1 {
			for a := 0; a < q+n+m+l+3; a++ {
				fmt.Print("-")
			}
			fmt.Print("\n")
		}
	}
	fmt.Printf("%d os, %d xs\n", os, xs)
}

func d2test(i, j int) bool {
	if i >= j {
		return false
	}
	return true
}

func fuse(vals ...int) string {
	var buf bytes.Buffer
	for _, v := range vals {
		fmt.Fprint(&buf, "%d", v)
	}
	return buf.String()
}

func d3test(i, j, k int) bool {
	// tosort := []int{i, j, k}
	// sort.Ints(tosort)
	// i = tosort[0]
	// j = tosort[1]
	// k = tosort[2]
	// key := fuse(i, j, k)
	// // if already seen skip
	// if d3map[key] {
	// 	return false
	// }
	// d3map[key] = true
	// diagonal
	// if i == j && i == k {
	// 	return true
	// }
	// two same
	if (i == j && i != k) || (i == k && i != j) || (j == k && i != j) {
		return true
	}
	// all different
	// if i != j && i != k && j != k {
	// 	return true
	// }
	return false
}

func d4test(i, j, k, l int) bool {
	tosort := []int{i, j, k, l}
	sort.Ints(tosort)
	i = tosort[0]
	j = tosort[1]
	k = tosort[2]
	l = tosort[3]
	key := fuse(i, j, k, l)
	// if already seen skip
	if d4map[key] {
		return false
	}
	d4map[key] = true
	// diagonal
	// if i == j && i == k && i == l {
	// 	return true
	// }
	// 3,1
	// switch {
	// case i == j && i == k && i != l:
	// 	return true
	// case i == j && i == l && i != k:
	// 	return true
	// case i == k && i == l && i != j:
	// 	return true
	// case j == k && j == l && j != i:
	// 	return true
	// }
	// 2,1,1
	// switch {
	// case i == j && i != k && i != l && k != l:
	// 	return true
	// case i == k && i != j && i != l && j != l:
	// 	return true
	// case i == l && i != j && i != k && j != k:
	// 	return true
	// case j == k && j != i && j != l && i != l:
	// 	return true
	// case j == l && j != i && j != k && i != k:
	// 	return true
	// case k == l && k != i && k != j && i != j:
	// 	return true
	// }
	// 2 and 2
	// switch {
	// case i == j && k == l && i != k:
	// 	return true
	// case i == k && j == l && i != j:
	// 	return true
	// case i == l && j == k && i != j:
	// 	return true
	// }
	// all different
	switch {
	case i != j && i != k && i != l && j != k && j != l && k != l:
		return true
	}
	return false
}

func d3last(n int) int {
	return (n*n*n - 3*n*n + 2*n) / 6
}

func d3formula(n int) int {
	return (4*n*n*n + 6*n*n + 2*n) / 3
}

func fact(n int) int {
	var tot int
	for tot = 1; n > 0; n-- {
		tot *= n
	}
	return tot
}

func d4211(n int) int {
	return n * (n - 1) * (n - 2) / 2
}

func d422(n int) int {
	// n > 3
	return n * (n - 1) / 2
}

func d41111(n int) int {
	return n * (n - 1) * (n - 2) * (n - 3) / 24
}

func d4formula(n int) int {
	return 2 * n * n * (n*n + 2) / 3
}

func total(n int) int {
	return 2*n*(n*n*n + 2*n*n + 8*n + 1)/3
}

func main() {
	n := 9
	// d2(5, 5, d2test)
	// d3(n, n, n, d3test)
	// fmt.Println(d3formula(9))
	// d4(n, n, n, n, d4test)
	// fmt.Println(d4formula(n))
	fmt.Println(total(n))
}
