package dynamic

import (
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	envelopsSort(envelopes)
	res := maxlis(envelopes)
	return len(res)
}
func envelopsSort(envelopes [][]int) {
	//先以w进行递增排序,h降序排列
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})
}

//寻找最长递增子序列
func maxlis(envelopes [][]int) []int {
	if len(envelopes) == 0 {
		return []int{}
	}
	dp := make([]int, 0)
	max := 0
	dp = append(dp, envelopes[0][1])
	for i := 1; i < len(envelopes); i++ {
		if dp[max] >= envelopes[i][1] {
			for j := 0; j <= max; j++ {
				if dp[j] >= envelopes[i][1] {
					dp[j] = envelopes[i][1]
					break
				}
			}
		} else {
			max++
			dp = append(dp, envelopes[i][1])
		}
	}
	return dp
}
