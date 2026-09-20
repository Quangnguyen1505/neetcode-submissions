func twoSum(nums []int, target int) []int {
    compare := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if j, ok := compare[complement]; ok {
			return []int{j, i}
		}
		compare[num] = i
	}
	return nil
}
