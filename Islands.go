package main

import "fmt"

func island_cnt(x [][]int8) int {
	cnt := 0
	for i := range x {
		for j := range x[i]{
			val := x[i][j]
			if val > 0 {
				if i == 0{
					if j == 0{
						cnt = cnt + 1
						}else{
						if x[i][j-1] == 0{
							cnt += 1
						}
					}
					}else{
					if j == 0{
						if x[i-1][j] == 0 {						
							cnt = cnt + 1
						}
						}else{
						if x[i][j-1] == 0 && x[i-1][j] == 0 {
							cnt += 1
						}
					}
				}
			}
		}
	}
	return cnt
}


func main() {
	Island1 := [][]int8{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	Island2 := [][]int8{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}
	fmt.Println(island_cnt(Island1))
	fmt.Println(island_cnt(Island2))
}
