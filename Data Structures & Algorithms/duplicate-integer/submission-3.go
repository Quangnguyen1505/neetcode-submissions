func hasDuplicate(nums []int) bool {
    dupMap := make(map[int]bool)
	for _, num := range nums { //O(n)
		if dupMap[num] { //O(1)
			return true
		}
		dupMap[num] = true 
	}
	return false
}

// TIME COMPLEXITY: O(n) -- SPACE COMPLEXITY: O(1)