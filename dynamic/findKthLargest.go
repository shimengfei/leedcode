package dynamic

//数组中的第K个最大元素
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) < k {
		return -1
	}
	//先进行排序
	quickSort(nums, 0, len(nums)-1)
	return nums[len(nums)-k]
}

//快速排序
func quickSort(nums []int, low int, high int) {
	if low < high {
		i, j := low, high
		//基线数据
		base := nums[low]
		for i < j {
			//右->左区间
			for i < j && nums[j] >= base {
				j--
			}
			if i < j {
				nums[i] = nums[j]
				i++
			}
			//左->右
			for i < j && nums[i] <= base {
				i++
			}
			if i < j {
				nums[j] = nums[i]
				j--
			}
		}
		nums[i] = base
		quickSort(nums, low, i-1)
		quickSort(nums, i+1, high)
	}
}
