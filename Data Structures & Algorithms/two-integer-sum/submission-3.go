// func twoSum(nums []int, target int) []int {
//     for i := 0; i < len(nums); i++ { // O(n)
//         for j := i + 1; j < len(nums); j++ { // O(n) => O(n) * O(n) = O(n^2)
//             if nums[i] + nums[j] == target {
//                 return []int{i,j}
//             }
//         }
//     }
//     return []int{}
// }
// ========== TIME COMPLEXITY: O(n^2) --- SPACE COMPLEXITY: O(1) ==============

// HASHMAP - MAP INX: VALUE, target - i = x, find x in map if true return, else false process continue 
func twoSum(nums []int, target int) []int {
    twoSumMap := make(map[int]int) 
    for i, num := range nums { // O(n)
        result := target - num 
        if j, ok := twoSumMap[result]; ok { // O(1)
            return []int{j, i}
        }
        twoSumMap[num] = i // O(1)
    }
    return []int{}
}

// ====== TIME COMPLEXITY: O(n) * O(1) = O(n) ---- SPACE COMPLEXITY: O(n) -> because insert new value of loop =====