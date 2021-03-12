package main

import "fmt"
// Solving x2 - Ny2 = 1.
// 使1729y^2 + 1成为完全平方数的最小正整数y是多少？
// 令1729y^2 + 1=n，n/y在40-43之间
// 正确答案： 446119244897052 – 1729×10728857123162 = 1, using 75 calculations.
// 网站链接：http://www.jakebakermaths.org.uk/maths/jshtmlpellsolverbigintegerv10.html

// 这个算法不行
func main() {

	for y := 1; y < 1e6; y++ {
		for n := 40*y + 1; n < 43*y; n++ {
			fmt.Printf("y: %v, n: %v\n", y, n)
			if 1729*y*y+1 == n*n {
				return
			}
		}
	}
}
