package dynamic

import "testing"

func TestEnvelopes(t *testing.T) {
	type envelopesTests struct {
		name      string
		envelopes [][]int
		result    int
		lis       [][]int
	}
	var tests []envelopesTests
	tests = append(tests, envelopesTests{
		name:      "base test",
		envelopes: [][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}},
		result:    3,
		lis:       [][]int{{2, 3}, {5, 4}, {6, 7}},
	})
	tests = append(tests, envelopesTests{
		name:      "base test",
		envelopes: [][]int{{5, 4}},
		result:    1,
		lis:       [][]int{{5, 4}},
	})
	tests = append(tests, envelopesTests{
		name:      "base test",
		envelopes: [][]int{},
		result:    0,
		lis:       [][]int{},
	})
	tests = append(tests, envelopesTests{
		name:      "base test",
		envelopes: [][]int{{1, 1}, {1, 1}, {1, 1}},
		result:    1,
		lis:       [][]int{},
	})
	for _, test := range tests {
		if res := maxEnvelopes(test.envelopes); res != test.result {
			t.Fatalf("error: %v  got: %v", test.result, res)
		}
	}
}

func TestMinPathSum(t *testing.T) {
	type envelopesTests struct {
		name   string
		grid   [][]int
		result int
		lis    []int
	}
	var tests []envelopesTests
	tests = append(tests, envelopesTests{
		name:   "base test",
		grid:   [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}},
		result: 7,
		lis:    []int{1, 3, 1, 1, 1},
	})
	tests = append(tests, envelopesTests{
		name:   "base test",
		grid:   [][]int{{1, 2, 3}, {4, 5, 6}},
		result: 12,
		lis:    []int{1, 3, 1, 1, 1},
	})
	for _, test := range tests {
		if res := minPathSum(test.grid); res != test.result {
			t.Fatalf("error: %v  got: %v", test.result, res)
		}
	}
}
