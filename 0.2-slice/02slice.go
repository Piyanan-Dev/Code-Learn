package main

import "fmt"

func main() {
	nums := []int{3, 8, 1, 9, 4}
	var total int
	var max = nums[0]
	var fourPlus int

	for _, i := range nums {
		total += i
		if max < i {
			max = i
		}
		if i > 4 {
			fourPlus += 1
		}
	}
	fmt.Println(total, max, fourPlus)
	nums = append(nums, 10)
	var total2 int
	for _, i := range nums {
		total2 += i
		if max < i {
			max = i
		}
	}
	fmt.Println(total2, max, fourPlus)

}
