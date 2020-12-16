package dynamic

//两数之和
//func twoSum(nums []int, target int) []int {
//	sort.Ints(nums)
//	length := len(nums)
//	right := length - 1
//	left := 0
//	result := make([]int, 0)
//	if length < 2 {
//		return result
//	}
//	for left < right {
//		if nums[left]+nums[right] == target {
//			result = append(result, nums[left], nums[right])
//			return result
//		}
//		if nums[left]+nums[right] < target {
//			for left < right && nums[left] == nums[left+1] {
//				left++
//			}
//			left++
//		}
//		if nums[left]+nums[right] > target {
//			for left < right && nums[right] == nums[right-1] {
//				right--
//			}
//			right--
//		}
//	}
//	return result
//}
func twoSum(nums []int, target int) []int {
	hash := map[int]int{}
	for i, num := range nums {
		if p, ok := hash[target-num]; ok {
			return []int{p, i}
		}
		hash[num] = i
	}
	return nil
}
