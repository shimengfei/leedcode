package dynamic

import "strings"

//无重复字符的最长子串
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	src := []rune(s)
	fast, low := 1, 0
	longest := 1
	for fast < len(src) {
		if index := strings.IndexRune(string(src[low:fast]), src[fast]); index != -1 {
			if fast-low > longest {
				longest = fast - low
			}
			low = low + index + 1
		}
		fast = fast + 1
	}
	if fast-low > longest {
		longest = fast - low
	}
	return longest
}
