package main

import "fmt"

/*
给定两个数组，编写一个函数来计算它们的交集。

说明：

输出结果中每个元素出现的次数，应与元素在两个数组中出现次数的最小值一致。
我们可以不考虑输出结果的顺序。
进阶：

如果给定的数组已经排好序呢？你将如何优化你的算法？
如果 nums1 的大小比 nums2 小很多，哪种方法更优？
如果 nums2 的元素存储在磁盘上，内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？

作者：力扣 (LeetCode)
链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-easy/x2y0c2/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

*/

func main() {
	nums1 := []int{61,24,20,58,95,53,17,32,45,85,70,20,83,62,35,89,5,95,12,86,58,77,30,64,46,13,5,92,67,40,20,38,31,18,89,85,7,30,67,34,62,35,47,98,3,41,53,26,66,40,54,44,57,46,70,60,4,63,82,42,65,59,17,98,29,72,1,96,82,66,98,6,92,31,43,81,88,60,10,55,66,82,0,79,11,81}
	nums2 := []int{5,25,4,39,57,49,93,79,7,8,49,89,2,7,73,88,45,15,34,92,84,38,85,34,16,6,99,0,2,36,68,52,73,50,77,44,61,48}

	//nums1 := []int{1, 2, 2, 3}
	//nums2 := []int{2, 3}
	ints := intersect(nums1, nums2)
	fmt.Println(ints)


	ints2 := []int{5,4,57,79,7,89,88,45,34,92,38,85,6,0,77,44,61}
	fmt.Println(intersect(ints, ints2))

}
// [5,4,57,79,7,89,88,45,34,92,38,85,6,0,77,44,61]

func intersect1(nums1 []int, nums2 []int) []int {
	ints := make([]int, 0)
	l2 := len(nums2)
	for _, i2 := range nums1 {
		for j := 0; j < len(nums2); j++ {
			if i2 == nums2[j] && len(ints) < len(nums1) && len(ints) < l2 {
				ints = append(ints, i2)
				if j == len(nums2) -1 {
					nums2 = nums2[:j]
 				} else {
					nums2 = append(nums2[:j], nums2[j+1:]...)
				}
				j--
				break
			}
		}
	}
	return ints
}

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	ints := make([]int, 0)

	for _, i := range nums1 {
		m[i]++
	}

	for _, i := range nums2 {
		times, ok := m[i]
		if ok && times > 0 {
			ints = append(ints, i)
			m[i]--
		}
	}

	return ints
}
