func hasDuplicate(nums []int) bool {
 	for i := 0; i < len(nums); i++ { // O(n)
		for j := i + 1; j < len(nums); j++ { // O(n)
			if nums[i] == nums[j] {
				return true
			}
		} 
 	}
	return false
}

// TIME COMLEXITY: O(n)*O(n) = O(n^2) --- SPACE COMPLEXITY: O(1)
