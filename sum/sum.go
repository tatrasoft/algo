package sum

import "fmt"

func Sum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

// recursive way of calling the function
func SumRecursivly(nums []int) int {
	fmt.Println(nums)
	if len(nums) == 0 {
		return 0
	}

	return nums[0] + SumRecursivly(nums[1:])
}