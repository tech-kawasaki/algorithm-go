package main

import "fmt"

// N個の正の整数a_0, a_1,..., a_N-1と正の整数Wが与えられます。
// a_0, a_1,..., a_N-1の中から何個かの整数を選んで総和をWとすることができるかどうか判定してください。
func main() {
	// N: 数列の数, W: 部分和
	var N, W int
	fmt.Scan(&N, &W)

	// a: 数列を格納するスライス
	a := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&a[i])
	}

	var flag = false

	// 数列の組み合わせの数だけループ
	for bit := 0; bit < 1<<N; bit++ {
		// 要素の和
		var sum int
		for i := 0; i < N; i++ {
			if bit&(1<<i) != 0 {
				sum += a[i]
			}
		}
		if sum == W {
			flag = true
		}
	}

	if flag {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
