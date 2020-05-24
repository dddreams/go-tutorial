package main

import "fmt"

func main() {
	nums := []int{2,7,11,15}
	target := 9
	var res = twoNum(nums, target)
	fmt.Println("结果：", res)
}

func twoNum(nums[] int, target int) []int{
	var len = len(nums)
	for i := 0; i < len; i++ {
		var m = make(map[int]int)
		for i = 0; i < len; i++ {
			var value int = m[i]
			if value != 0 {
				return []int{value, i}
			}
			m[target - nums[i]] = i
		}
	}	
	return nil
}
