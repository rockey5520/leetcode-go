package main

import "log"

func main() {
	log.Println(twoSumOptimized([]int{3, 2, 4}, 6))
}

func twoSumOptimized(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		v, found := m[target-num]
		if found {
			return []int{i, v}
		}
		m[num] = i
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	var pos_x int
	var pos_y int
	var isTargetReached bool
	for i, num := range nums {
		log.Println(i, num)
		for j := i; j < len(nums)-1; j++ {
			log.Println(nums[i], nums[j+1])
			if nums[i]+nums[j+1] == target {
				pos_x = i
				pos_y = j + 1
				isTargetReached = true
				break
			}
		}
		if isTargetReached {
			break
		}
	}
	return []int{pos_x, pos_y}
}
