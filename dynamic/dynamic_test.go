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

func TestLemonadeChange(t *testing.T) {
	type lemonadeChangeTests struct {
		name   string
		grid   []int
		result bool
		lis    []int
	}
	var tests []lemonadeChangeTests
	tests = append(tests, lemonadeChangeTests{
		name:   "base test",
		grid:   []int{},
		result: true,
	})
	tests = append(tests, lemonadeChangeTests{
		name:   "base test",
		grid:   []int{5, 5, 5, 10, 20},
		result: true,
	})
	tests = append(tests, lemonadeChangeTests{
		name:   "base test",
		grid:   []int{10, 10},
		result: false,
	})
	tests = append(tests, lemonadeChangeTests{
		name:   "base test",
		grid:   []int{5, 5, 10, 10, 20},
		result: false,
	})
	tests = append(tests, lemonadeChangeTests{
		name:   "base test",
		grid:   []int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 10, 5, 20, 20, 5, 5, 5, 5, 5, 10, 5, 5, 5, 20, 5, 5, 5, 10, 5, 5, 10, 5, 20, 5, 5, 20, 5, 10, 5, 5, 20, 5, 5, 5, 5, 5, 5, 10, 20, 5, 20, 20, 10, 5, 20, 20, 5, 10, 5, 5, 5, 5, 5, 5, 20, 20, 20, 20, 5, 5, 10, 5, 20, 5, 5, 5, 5, 10, 10, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 20, 5, 20, 10, 10, 20, 5, 5, 5, 5, 20, 20, 5, 5, 5, 5, 20, 5, 20, 20, 5, 5, 10, 5, 5, 5, 20, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 10, 20, 20, 5, 5, 10, 20, 5, 5, 5, 5, 10, 20, 5, 5, 10, 20, 5, 10, 10, 5, 5, 5, 5, 5, 5, 10, 10, 10, 5, 10, 5, 10, 5, 5, 10, 10, 5, 5, 5, 20, 5, 20, 10, 20, 5, 5, 5, 20, 10, 5, 5, 20, 5, 5, 5, 10, 5, 5, 10, 5, 5, 20, 5, 10, 10, 5, 5, 10, 5, 5, 10, 5, 10, 5, 20, 10, 5, 10, 10, 5, 5, 5, 5, 10, 5, 5, 5, 20, 5, 5, 5, 5, 10, 5, 10, 10, 5, 20, 20, 5, 10, 10, 10, 5, 10, 5, 10, 5, 10, 5, 5, 10, 5, 5, 5, 20, 5, 5, 20, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 20, 5, 5, 5, 5, 10, 5, 5, 5, 20, 5, 5, 5, 5, 10, 20, 5, 5, 5, 20, 20, 5, 10, 5, 5, 5, 10, 5, 10, 20, 20, 5, 5, 5, 5, 5, 5, 20, 10, 5, 10, 5, 5, 20, 10, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 5, 10, 5, 10, 20, 10, 10, 5, 5, 20, 10, 20, 5, 5, 5, 10, 5, 5, 5, 10, 5, 20, 20, 20, 10, 20, 5, 5, 5, 5, 20, 5, 20, 5, 10, 5, 5, 5, 5, 5, 5, 20, 5, 10, 5, 5, 5, 20, 5, 5, 5, 10, 10, 5, 5, 5, 5, 5, 20, 20, 20, 5, 5, 5, 5, 20, 5, 20, 5, 20, 20, 10, 10, 5, 5, 5, 20, 10, 20, 10, 20, 5, 20, 5, 5, 5, 10, 10, 5, 5, 5, 5, 10, 10, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 10, 20, 20, 5, 20, 10, 5, 5, 20, 5, 5, 10, 5, 5, 5, 5, 10, 5, 5, 5, 20, 5, 10, 5, 10, 10, 20, 5, 20, 5, 20, 10, 5, 5, 5, 20, 5, 10, 10, 5, 5, 10, 5, 10, 5, 5, 20, 20, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 10, 5, 10, 20, 20, 5, 5, 20, 5, 10, 5, 20, 5, 20, 20, 5, 5, 5, 20, 20, 20, 5, 5, 5, 5, 20, 10, 5, 5, 10, 10, 10, 5, 10, 5, 10, 5, 20, 5, 5, 5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 20, 20, 10, 10, 5, 5, 20, 5, 5, 5, 5, 20, 20, 20, 20, 5, 5, 20, 20, 5, 20, 5, 5, 5, 10, 20, 10, 5, 20, 5, 5, 5, 5, 10, 10, 5, 10, 5, 5, 10, 5, 5, 20, 10, 20, 5, 5, 5, 10, 5, 5, 10, 10, 5, 20, 5, 5, 20, 5, 5, 20, 10, 10, 5, 5, 10, 5, 5, 20, 5, 10, 5, 20, 20, 10, 10, 20, 5, 5, 5, 20, 5, 5, 20, 20, 10, 20, 10, 10, 5, 20, 10, 5, 10, 5, 10, 5, 5, 20, 5, 20, 20, 5, 5, 5, 5, 20, 10, 5, 5, 5, 5, 5, 20, 5, 5, 20, 10, 5, 5, 5, 5, 5, 5, 10, 5, 10, 5, 20, 20, 20, 20, 5, 5, 20, 5, 5, 5, 20, 5, 20, 5, 5, 5, 10, 5, 20, 5, 5, 5, 5, 20, 5, 20, 20, 5, 5, 5, 5, 10, 5, 20, 20, 5, 20, 5, 10, 5, 5, 5, 5, 20, 5, 5, 5, 20, 20, 5, 5, 5, 20, 10, 20, 5, 5, 10, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 20, 10, 20, 20, 5, 5, 10, 10, 5, 10, 10, 5, 5, 10, 10, 5, 5, 5, 5, 5, 10, 5, 20, 5, 10, 5, 20, 10, 5, 10, 5, 10, 20, 10, 5, 5, 5, 20, 5, 5, 20, 5, 5, 5, 5, 5, 10, 10, 20, 20, 20, 5, 20, 20, 5, 20, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 10, 5, 5, 5, 5, 10, 20, 20, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 20, 20, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 20, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 10, 5, 10, 5, 5, 5, 20, 5, 5, 5, 20, 5, 20, 5, 20, 10, 5, 5, 5, 5, 10, 5, 10, 5, 20, 20, 10, 5, 5, 20, 10, 10, 5, 10, 20, 5, 5, 5, 10, 5, 20, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 20, 5, 5, 10, 10, 5, 5, 10, 5, 10, 10, 20, 10, 20, 5, 5, 5, 10, 10, 5, 5, 20, 5, 20, 5, 5, 5, 5, 10, 5, 20, 5, 5, 5, 20, 5, 5, 10, 10, 5, 5, 5, 20, 5, 5, 10, 5, 5, 5, 5, 20, 5, 10, 5, 5, 10, 5, 10, 10, 5, 10, 10, 5, 20, 20, 5, 5, 20, 5, 5, 5, 5, 5, 20, 10, 10, 5, 10, 5, 20, 20, 5, 5, 5, 5, 5, 10, 5, 20, 10, 20, 5, 20, 5, 20, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 10, 5, 5, 20, 20, 5, 20, 20, 5, 5, 5, 10, 20, 5, 5, 10, 10, 5, 5, 20, 10, 20, 5, 10, 5, 5, 5, 5, 20, 5, 5, 5, 5, 20, 20, 5, 20, 5, 10, 10, 5, 10, 20, 20, 20, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 20, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 20, 10, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 20, 10, 5, 5, 5, 5, 5, 10, 5, 20, 5, 5, 5, 5, 10, 5, 10, 20, 10, 20, 5, 5, 5, 5, 10, 5, 20, 5, 20, 5, 20, 5, 10, 5, 10, 10, 10, 5, 5, 5, 20, 20, 5, 5, 10, 10, 10, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 10, 5, 5, 5, 5, 5, 5, 5, 20, 10, 10, 20, 10, 5, 5, 5, 5, 5, 5, 20, 5, 5, 20, 5, 10, 5, 20, 5, 10, 5, 5, 20, 5, 5, 5, 10, 5, 5, 5, 10, 5, 5, 20, 20, 5, 5, 10, 20, 10, 20, 5, 10, 20, 5, 5, 10, 5, 20, 20, 5, 5, 5, 5, 5, 5, 5, 10, 10, 10, 5, 20, 20, 5, 5, 20, 20, 10, 20, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 20, 10, 5, 5, 5, 5, 10, 5, 10, 5, 5, 5, 5, 10, 20, 20, 5, 5, 5, 5, 20, 5, 5, 5, 5, 10, 20, 5, 5, 10, 5, 5, 10, 5, 10, 20, 20, 5, 10, 5, 10, 5, 5, 10, 10, 5, 5, 5, 5, 5, 5, 20, 5, 20, 20, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 20, 10, 10, 5, 5, 10, 5, 20, 5, 5, 10, 10, 20, 5, 5, 5, 20, 10, 5, 5, 5, 10, 5, 5, 5, 5, 5, 20, 5, 10, 5, 5, 10, 20, 5, 5, 20, 10, 5, 5, 20, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 10, 10, 5, 5, 5, 5, 10, 5, 20, 5, 10, 10, 5, 10, 5, 5, 5, 5, 5, 5, 10, 20, 5, 5, 20, 20, 5, 20, 10, 5, 5, 5, 10, 5, 20, 5, 5, 10, 20, 5, 5, 20, 10, 20, 10, 5, 5, 5, 10, 5, 5, 5, 5, 5, 10, 20, 10, 5, 5, 5, 5, 5, 5, 5, 5, 20, 10, 5, 20, 10, 5, 20, 10, 20, 10, 5, 20, 5, 5, 5, 5, 5, 20, 20, 10, 10, 10, 10, 5, 5, 20, 5, 10, 5, 5, 10, 20, 5, 5, 5, 5, 20, 5, 5, 5, 10, 5, 5, 10, 5, 20, 20, 10, 10, 5, 5, 10, 5, 5, 10, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 10, 5, 5, 5, 5, 10, 5, 5, 20, 10, 10, 5, 20, 20, 20, 5, 5, 5, 5, 20, 5, 10, 5, 5, 5, 5, 5, 10, 20, 5, 20, 5, 20, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 10, 5, 20, 5, 5, 20, 5, 5, 5, 5, 5, 20, 10, 5, 10, 10, 5, 20, 10, 5, 20, 20, 5, 5, 20, 5, 5, 10, 5, 5, 10, 5, 5, 10, 10, 5, 5, 20, 10, 5, 20, 5, 10, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 20, 5, 10, 5, 10, 20, 10, 10, 5, 5, 10, 5, 20, 10, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 10, 20, 20, 5, 5, 5, 20, 5, 5, 5, 5, 10, 5, 20, 5, 5, 5, 20, 10, 5, 5, 5, 5, 10, 5, 5, 5, 10, 5, 5, 5, 5, 20, 5, 5, 5, 10, 20, 20, 20, 5, 20, 5, 10, 5, 10, 20, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 20, 10, 5, 20, 5, 5, 5, 5, 5, 5, 10, 5, 5, 10, 5, 5, 5, 5, 20, 5, 5, 5, 20, 5, 5, 10, 5, 10, 5, 10, 20, 5, 5, 5, 5, 20, 10, 5, 5, 5, 20, 5, 10, 10, 5, 5, 10, 20, 5, 10, 5, 5, 5, 20, 10, 5, 20, 5, 5, 5, 20, 20, 10, 20, 10, 5, 20, 5, 10, 10, 20, 10, 5, 5, 10, 5, 20, 20, 5, 5, 5, 5, 10, 5, 10, 20, 20, 20, 10, 20, 5, 10, 5, 5, 5, 10, 5, 5, 5, 5, 10, 5, 5, 5, 20, 10, 20, 5, 20, 5, 5, 5, 20, 10, 5, 5, 20, 5, 10, 5, 5, 5, 10, 5, 5, 10, 10, 5, 20, 20, 20, 5, 5, 5, 5, 20, 5, 20, 20, 5, 10, 5, 20, 5, 10, 20, 5, 5, 10, 5, 5, 5, 5, 5, 10, 10, 5, 5, 5, 20, 5, 10, 5, 5, 5, 5, 5, 20, 5, 10, 10, 5, 5, 10, 20, 10, 5, 5, 5, 5, 10, 5, 5, 10, 5, 10, 5, 10, 5, 20, 5, 5, 10, 5, 20, 5, 5, 10, 5, 20, 5, 20, 20, 10, 5, 20, 10, 20, 20, 5, 20, 5, 20, 5, 5, 20, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 20, 20, 10, 5, 5, 20, 20, 20, 10, 20, 20, 20, 10, 10, 20, 5, 10, 5, 5, 10, 5, 5, 5, 20, 10, 20, 20, 20, 5, 5, 5, 10, 5, 5, 5, 20, 10, 5, 10, 5, 5, 10, 5, 5, 5, 20, 5, 20, 5, 5, 5, 5, 5, 10, 20, 10, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 10, 20, 5, 10, 20, 10, 10, 5, 5, 20, 5, 20, 5, 5, 20, 5, 20, 20, 5, 5, 5, 10, 20, 5, 5, 10, 5, 5, 5, 5, 5, 10, 5, 5, 5, 20, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 10, 5, 20, 10, 5, 5, 10, 5, 5, 10, 10, 10, 5, 5, 10, 5, 10, 5, 20, 10, 20, 10, 20, 20, 10, 5, 10, 10, 10, 5, 20, 5, 5, 5, 5, 10, 5, 5, 20, 20, 10, 10, 10, 5, 10, 5, 5, 20, 5, 5, 5, 5, 5, 5, 20, 5, 5, 10, 20, 5, 5, 5, 20, 10, 10, 10, 20, 20, 10, 5, 5, 10, 20, 20, 10, 5, 5, 5, 10, 10, 5, 20, 5, 10, 20, 20, 5, 5, 20, 10, 5, 20, 5, 5, 10, 5, 5, 5, 10, 5, 5, 10, 5, 5, 10, 5, 20, 5, 5, 20, 5, 5, 5, 20, 20, 5, 20, 20, 5, 5, 10, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 10, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 20, 20, 5, 10, 5, 5, 5, 5, 5, 10, 5, 20, 10, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 20, 5, 10, 5, 5, 20, 5, 5, 20, 20, 20, 20, 5, 10, 5, 20, 5, 5, 20, 5, 5, 5, 10, 5, 10, 5, 20, 5, 5, 5, 5, 20, 5, 5, 20, 20, 5, 5, 20, 5, 5, 5, 20, 10, 5, 5, 5, 5, 5, 20, 20, 5, 5, 5, 5, 5, 10, 10, 5, 10, 10, 10, 5, 5, 5, 20, 5, 20, 10, 20, 20, 5, 5, 10, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 10, 10, 5, 10, 10, 5, 10, 5, 20, 10, 10, 5, 5, 20, 5, 5, 10, 10, 10, 5, 10, 10, 10, 5, 10, 5, 5, 10, 5, 10, 20, 20, 5, 20, 5, 5, 5, 10, 10, 20, 5, 5, 5, 10, 5, 5, 5, 5, 5, 20, 5, 5, 10, 20, 5, 20, 5, 5, 10, 20, 10, 20, 5, 10, 5, 5, 5, 10, 20, 10, 5, 10, 5, 5, 5, 20, 5, 5, 10, 5, 5, 5, 5, 5, 20, 10, 10, 5, 20, 5, 20, 5, 5, 10, 5, 5, 20, 5, 20, 20, 10, 20, 10, 5, 20, 5, 20, 20, 20, 5, 5, 5, 10, 5, 20, 20, 10, 5, 20, 5, 5, 5, 20, 10, 5, 5, 20, 20, 5, 20, 20, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 5, 5, 5, 20, 5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 20, 5, 5, 5, 5, 5, 20, 5, 5, 10, 10, 20, 5, 20, 5, 5, 5, 20, 10, 5, 5, 5, 5, 10, 5, 20, 5, 5, 5, 5, 20, 20, 10, 5, 5, 5, 10, 5, 20, 5, 10, 5, 20, 5, 10, 5, 5, 5, 5, 5, 20, 20, 5, 5, 5, 5, 5, 20, 10, 20, 5, 5, 5, 10, 5, 5, 5, 20, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5, 5, 20, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 5, 10, 5, 10, 5, 5, 10, 5, 5, 10, 10, 20, 5, 20, 5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 5, 10, 5, 5, 10, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 20, 5, 5, 5, 20, 10, 5, 5, 10, 5, 10, 5, 20, 5, 10, 5, 5, 20, 5, 10, 5, 5, 5, 10, 5, 5, 5, 5, 5, 20, 5, 5, 5, 20, 5, 10, 5, 20, 5, 5, 5, 10, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 5, 20, 20, 5, 5, 10, 5, 5, 5, 5, 10, 10, 10, 5, 5, 20, 5, 10, 20, 5, 5, 5, 5, 5, 10, 10, 5, 20, 5, 10, 5, 5, 5, 5, 20, 5, 20, 10, 20, 10, 10, 20, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 5, 10, 5, 5, 5, 20, 20, 20, 5, 10, 20, 20, 10, 20, 20, 20, 5, 10, 20, 5, 10, 5, 10, 5, 20, 20, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 20, 10, 5, 5, 20, 5, 5, 5, 20, 5, 10, 20, 10, 20, 5, 20, 5, 5, 5, 5, 20, 5, 5, 5, 5, 20, 5, 5, 5, 5, 20, 5, 5, 20, 5, 20, 10, 5, 5, 5, 20, 10, 5, 5, 5, 20, 5, 10, 5, 5, 5, 5, 10, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 10, 5, 5, 5, 10, 10, 5, 5, 10, 20, 20, 20, 20, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 5, 5, 20, 5, 20, 5, 10, 20, 5, 10, 20, 5, 20, 5, 10, 20, 5, 20, 5, 20, 20, 5, 10, 5, 20, 5, 5, 5, 20, 10, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 20, 5, 5, 5, 5, 5, 5, 5, 10, 20, 10, 10, 10, 20, 5, 5, 10, 20, 5, 20, 20, 10, 5, 5, 20, 5, 20, 5, 20, 10, 5, 10, 5, 5, 5, 5, 5, 10, 5, 5, 10, 5, 5, 20, 5, 20, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 10, 5, 10, 20, 5, 20, 5, 10, 10, 5, 5, 10, 5, 20, 20, 5, 5, 5, 10, 20, 5, 5, 5, 5, 10, 5, 5, 5, 5, 20, 5, 10, 10, 5, 5, 20, 5, 10, 5, 20, 10, 10, 5, 10, 20, 5, 5, 10, 5, 5, 10, 5, 5, 5, 5, 10, 5, 5, 5, 10, 5, 10, 20, 20, 5, 10, 5, 20, 5, 10, 5, 5, 20, 5, 5, 20, 5, 5, 20, 5, 5, 10, 20, 5, 5, 5, 20, 20, 20, 5, 20, 10, 5, 5, 5, 5, 5, 5, 10, 5, 20, 5, 10, 5, 20, 5, 20, 5, 20, 5, 10, 5, 20, 20, 10, 5, 5, 10, 20, 5, 5, 20, 10, 5, 5, 5, 5, 5, 5, 5, 5, 20, 10, 10, 10, 10, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 20, 20, 5, 5, 5, 5, 5, 5, 10, 5, 20, 20, 10, 20, 5, 5, 5, 5, 10, 20, 5, 20, 5, 5, 20, 5, 5, 5, 5, 5, 20, 20, 5, 5, 5, 10, 5, 5, 10, 10, 10, 20, 10, 5, 20, 5, 5, 10, 20, 5, 5, 10, 5, 5, 20, 5, 20, 5, 5, 5, 5, 5, 5, 20, 5, 20, 5, 5, 20, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 5, 5, 5, 10, 20, 10, 5, 20, 20, 5, 20, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 10, 10, 20, 10, 5, 5, 10, 5, 5, 10, 10, 5, 5, 5, 5, 5, 20, 10, 5, 5, 20, 5, 5, 20, 5, 10, 10, 20, 20, 20, 10, 20, 5, 20, 20, 5, 5, 20, 5, 20, 5, 5, 20, 5, 20, 5, 5, 10, 5, 20, 10, 5, 5, 5, 5, 10, 5, 20, 5, 20, 20, 20, 5, 10, 5, 5, 5, 5, 10, 20, 10, 10, 20, 5, 10, 5, 5, 10, 5, 20, 5, 20, 10, 20, 5, 5, 5, 5, 20, 20, 5, 5, 5, 10, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 20, 5, 5, 5, 20, 5, 10, 20, 5, 20, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 10, 5, 10, 20, 5, 5, 5, 5, 5, 5, 10, 5, 10, 5, 10, 5, 5, 5, 10, 10, 20, 5, 5, 10, 20, 5, 20, 5, 20, 10, 5, 5, 5, 20, 5, 10, 5, 5, 5, 5, 5, 20, 20, 10, 5, 5, 10, 5, 5, 5, 10, 5, 10, 10, 5, 5, 5, 10, 5, 10, 5, 5, 5, 5, 5, 10, 5, 10, 10, 10, 10, 5, 5, 5, 5, 5, 20, 10, 5, 5, 5, 10, 5, 5, 20, 5, 5, 10, 5, 5, 10, 10, 10, 5, 10, 20, 5, 5, 20, 5, 20, 5, 5, 20, 5, 5, 20, 20, 5, 5, 20, 10, 5, 5, 5, 5, 5, 5, 5, 10, 5, 20, 5, 10, 5, 5, 5, 5, 5, 5, 5, 10, 5, 20, 10, 5, 20, 5, 5, 5, 5, 5, 5, 20, 5, 10, 5, 5, 10, 20, 20, 5, 5, 20, 5, 5, 20, 5, 10, 20, 20, 5, 5, 5, 5, 5, 5, 20, 20, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 20, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 20, 20, 20, 5, 5, 5, 10, 10, 5, 5, 10, 5, 5, 5, 10, 10, 5, 5, 10, 10, 10, 20, 5, 20, 5, 5, 5, 5, 5, 5, 5, 10, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 10, 10, 5, 10, 5, 5, 5, 5, 10, 10, 20, 10, 20, 5, 10, 20, 20, 20, 5, 5, 10, 5, 20, 20, 5, 20, 20, 5, 10, 5, 5, 10, 20, 5, 5, 5, 5, 5, 10, 5, 10, 5, 10, 20, 5, 5, 5, 5, 5, 20, 10, 20, 10, 20, 20, 5, 5, 20, 20, 5, 5, 10, 5, 5, 5, 20, 10, 10, 5, 5, 20, 5, 5, 5, 20, 10, 10, 5, 5, 5, 10, 5, 10, 10, 20, 5, 10, 5, 20, 5, 5, 5, 20, 10, 20, 5, 10, 5, 20, 10, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 20, 5, 20, 5, 5, 20, 5, 10, 5, 5, 5, 5, 5, 5, 10, 5, 10, 5, 5, 20, 5, 10, 5, 5, 5, 10, 10, 5, 5, 5, 20, 5, 5, 20, 20, 5, 5, 10, 20, 5, 5, 5, 5, 5, 10, 20, 5, 5, 10, 20, 10, 20, 5, 10, 5, 5, 20, 20, 5, 5, 5, 10, 5, 5, 10, 5, 5, 5, 20, 20, 10, 5, 5, 20, 20, 5, 20, 5, 10, 20, 5, 5, 5, 10, 5, 5, 5, 10, 5, 5, 20, 10, 5, 20, 20, 20, 5, 5, 10, 5, 20, 20, 20, 20, 10, 5, 10, 10, 10, 5, 5, 5, 20, 5, 10, 10, 20, 5, 5, 5, 5, 5, 20, 5, 5, 20, 10, 5, 5, 20, 10, 5, 10, 10, 20, 10, 20, 5, 5, 5, 10, 10, 5, 5, 10, 20, 20, 5, 5, 5, 10, 5, 5, 20, 5, 5, 5, 5, 5, 10, 20, 10, 5, 10, 5, 5, 5, 5, 5, 5, 10, 5, 5, 10, 10, 5, 20, 5, 10, 5, 5, 5, 10, 5, 5, 10, 20, 10, 10, 20, 10, 10, 20, 5, 5, 5, 5, 10, 5, 5, 20, 20, 5, 10, 20, 10, 5, 5, 20, 5, 20, 20, 5, 5, 5, 10, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 5, 5, 5, 20, 5, 5, 5, 10, 20, 5, 5, 5, 5, 10, 5, 10, 20, 5, 10, 5, 5, 5, 20, 5, 20, 10, 5, 5, 5, 5, 20, 5, 5, 5, 10, 5, 5, 5, 5, 10, 20, 5, 5, 5, 10, 20, 10, 10, 5, 5, 20, 20, 20, 10, 20, 20, 10, 10, 5, 5, 5, 10, 5, 20, 20, 5, 5, 20, 20, 5, 5, 20, 5, 20, 5, 5, 5, 5, 10, 5, 10, 5, 5, 10, 5, 20, 5, 5, 10, 10, 5, 10, 5, 5, 5, 5, 5, 10, 5, 5, 10, 10, 5, 5, 5, 5, 5, 20, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 10, 20, 10, 5, 5, 10, 5, 5, 5, 5, 5, 10, 20, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 5, 5, 10, 5, 10, 20, 20, 10, 5, 5, 20, 5, 10, 5, 5, 5, 5, 20, 5, 10, 20, 5, 5, 10, 20, 5, 5, 10, 10, 20, 5, 5, 5, 5, 5, 20, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 20, 20, 5, 5, 5, 5, 5, 5, 10, 10, 20, 5, 5, 20, 10, 5, 5, 20, 20, 20, 20, 20, 5, 20, 10, 5, 20, 5, 5, 5, 5, 10, 5, 20, 5, 5, 10, 5, 5, 10, 10, 20, 5, 5, 20, 5, 20, 20, 5, 5, 5, 5, 5, 5, 5, 10, 20, 10, 20, 10, 5, 5, 5, 10, 5, 20, 20, 10, 10, 20, 20, 5, 5, 20, 5, 10, 10, 5, 5, 20, 5, 20, 5, 5, 5, 10, 5, 5, 20, 5, 10, 5, 10, 5, 10, 5, 10, 10, 5, 20, 5, 20, 5, 10, 5, 5, 5, 5, 20, 5, 20, 5, 5, 20, 10, 5, 10, 10, 5, 5, 5, 5, 10, 5, 10, 5, 10, 5, 10, 20, 5, 5, 10, 5, 10, 5, 10, 10, 20, 20, 20, 20, 10, 5, 5, 5, 20, 5, 5, 10, 10, 20, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 20, 10, 10, 10, 10, 20, 10, 5, 5, 5, 5, 10, 20, 5, 5, 5, 5, 20, 5, 5, 5, 10, 5, 5, 10, 5, 10, 10, 5, 5, 5, 5, 20, 20, 20, 5, 20, 5, 5, 10, 10, 20, 5, 5, 5, 5, 20, 5, 20, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 20, 20, 10, 20, 5, 5, 5, 20, 10, 5, 10, 10, 20, 20, 5, 20, 10, 5, 5, 5, 20, 5, 5, 10, 5, 5, 5, 20, 10, 5, 5, 5, 10, 20, 5, 20, 20, 5, 20, 5, 5, 10, 10, 10, 5, 5, 20, 20, 5, 5, 20, 5, 5, 10, 20, 10, 20, 5, 5, 20, 5, 10, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 20, 5, 20, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 5, 20, 5, 5, 20, 5, 10, 5, 20, 20, 5, 10, 5, 5, 5, 5, 10, 5, 5, 5, 5, 10, 5, 5, 5, 10, 10, 5, 20, 20, 5, 5, 5, 20, 10, 5, 5, 5, 5, 5, 10, 20, 20, 20, 10, 10, 5, 5, 20, 10, 5, 10, 5, 5, 5, 5, 5, 20, 5, 20, 20, 5, 5, 5, 10, 5, 10, 10, 10, 5, 5, 10, 5, 5, 5, 5, 5, 5, 20, 5, 5, 10, 5, 5, 5, 5, 5, 5, 10, 5, 20, 20, 5, 5, 5, 20, 5, 5, 5, 5, 10, 5, 5, 10, 5, 20, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 5, 10, 5, 5, 5, 20, 5, 10, 5, 5, 10, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 5, 20, 20, 5, 5, 5, 5, 5, 10, 10, 20, 5, 5, 10, 5, 10, 20, 10, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 20, 20, 5, 5, 5, 5, 5, 20, 5, 5, 20, 5, 10, 20, 5, 5, 5, 10, 5, 10, 5, 20, 10, 5, 10, 5, 5, 20, 20, 5, 10, 5, 5, 5, 10, 5, 5, 20, 10, 10, 20, 5, 20, 5, 5, 20, 5, 5, 10, 20, 5, 10, 20, 5, 5, 5, 5, 5, 5, 20, 5, 10, 5, 5, 20, 5, 5, 20, 5, 10, 5, 10, 10, 5, 5, 20, 5, 5, 10, 5, 5, 5, 10, 5, 5, 5, 5, 20, 5, 20, 10, 5, 5, 20, 5, 20, 5, 5, 5, 5, 5, 5, 20, 5, 20, 20, 5, 10, 10, 20, 20, 5, 5, 5, 5, 5, 20, 20, 20, 5, 10, 5, 5, 10, 5, 20, 10, 10, 20, 5, 5, 10, 5, 5, 10, 5, 5, 5, 20, 5, 5, 20, 5, 5, 5, 5, 20, 10, 5, 10, 10, 10, 10, 10, 5, 10, 5, 5, 5, 20, 10, 20, 10, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 10, 5, 5, 5, 5, 5, 20, 5, 20, 5, 5, 20, 5, 5, 20, 5, 5, 5, 20, 10, 20, 5, 5, 10, 5, 5, 5, 20, 20, 5, 10, 5, 5, 20, 20, 5, 5, 5, 10, 5, 10, 5, 5, 5, 20, 10, 10, 5, 5, 5, 5, 5, 5, 5, 20, 20, 5, 20, 5, 10, 5, 20, 10, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 20, 20, 10, 10, 5, 5, 20, 5, 10, 5, 5, 5, 10, 5, 20, 5, 5, 5, 5, 5, 5, 5, 10, 5, 10, 5, 5, 5, 5, 20, 20, 5, 20, 5, 20, 5, 20, 5, 5, 20, 5, 5, 10, 20, 5, 20, 5, 20, 5, 5, 5, 10, 10, 5, 5, 20, 5, 5, 10, 5, 5, 10, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 10, 10, 5, 5, 5, 5, 10, 5, 10, 10, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 10, 5, 5, 5, 10, 5, 5, 5, 10, 5, 20, 5, 20, 10, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 10, 20, 5, 10, 5, 5, 20, 10, 5, 5, 5, 10, 10, 10, 5, 5, 5, 5, 5, 5, 20, 10, 20, 20, 5, 5, 20, 5, 5, 5, 5, 20, 5, 10, 5, 10, 5, 5, 20, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 20, 10, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 10, 10, 5, 10, 5, 5, 20, 5, 5, 10, 5, 5, 10, 20, 10, 5, 20, 5, 5, 5, 20, 20, 5, 10, 5, 5, 5, 5, 10, 10, 10, 20, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 20, 5, 10, 5, 5, 5, 5, 10, 5, 5, 5, 20, 5, 5, 20, 5, 5, 5, 5, 5, 10, 20, 5, 20, 5, 10, 10, 5, 5, 20, 5, 5, 5, 5, 5, 20, 5, 5, 5, 20, 20, 5, 20, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 5, 5, 10, 5, 5, 5, 10, 20, 20, 5, 5, 5, 5, 5, 10, 5, 20, 20, 5, 5, 5, 20, 10, 5, 5, 5, 10, 5, 5, 5, 5, 20, 20, 5, 20, 10, 10, 20, 5, 5, 20, 5, 10, 20, 20, 10, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 20, 5, 5, 5, 20, 5, 20, 5, 5, 5, 10, 10, 20, 5, 5, 5, 5, 10, 10, 5, 20, 5, 20, 5, 5, 5, 5, 20, 5, 20, 10, 5, 20, 5, 10, 10, 5, 20, 5, 5, 5, 10, 20, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 20, 20, 5, 20, 5, 5, 10, 20, 5, 10, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 10, 5, 20, 5, 5, 5, 10, 10, 20, 5, 5, 5, 5, 5, 5, 20, 5, 5, 10, 20, 20, 5, 10, 20, 5, 5, 20, 10, 20, 5, 10, 5, 5, 20, 10, 5, 5, 20, 10, 5, 10, 10, 5, 5, 20, 20, 5, 20, 10, 5, 5, 5, 5, 5, 20, 5, 5, 10, 10, 10, 10, 5, 5, 5, 5, 5, 5, 5, 20, 20, 5, 20, 5, 10, 20, 5, 10, 5, 5, 5, 10, 5, 20, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 10, 20, 5, 5, 10, 5, 20, 20, 10, 10, 5, 5, 5, 5, 10, 20, 5, 10, 5, 20, 10, 5, 5, 20, 5, 20, 10, 5, 5, 5, 5, 5, 10, 5, 5, 5, 20, 5, 20, 5, 20, 5, 5, 20, 10, 20, 5, 20, 5, 20, 5, 10, 20, 20, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 20, 5, 10, 5, 5, 20, 5, 5, 10, 5, 5, 5, 5, 20, 5, 10, 5, 5, 5, 5, 5, 5, 5, 10, 10, 5, 20, 20, 5, 5, 5, 5, 20, 5, 20, 5, 5, 5, 5, 5, 10, 10, 20, 20, 5, 5, 5, 20, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 10, 20, 5, 5, 5, 20, 5, 10, 5, 5, 5, 5, 5, 20, 5, 5, 5, 10, 20, 20, 5, 10, 10, 5, 5, 10, 5, 5, 20, 10, 10, 5, 5, 10, 5, 5, 20, 5, 5, 5, 5, 20, 5, 10, 10, 5, 5, 10, 20, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 10, 5, 10, 10, 10, 5, 20, 10, 5, 5, 5, 5, 5, 20, 5, 5, 5, 10, 20, 5, 5, 5, 5, 5, 10, 5, 5, 5, 10, 5, 20, 20, 5, 20, 20, 20, 5, 10, 5, 5, 5, 5, 20, 20, 5, 5, 5, 10, 5, 5, 5, 10, 10, 5, 5, 5, 5, 5, 5, 5, 10, 5, 10, 5, 5, 5, 5, 20, 5, 20, 5, 20, 5, 5, 20, 5, 5, 5, 20, 10, 5, 5, 5, 5, 5, 5, 5, 10, 5, 20, 5, 5, 5, 5, 20, 20, 5, 20, 5, 10, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 20, 5, 20, 10, 10, 5, 5, 5, 10, 5, 20, 5, 20, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 20, 10, 5, 10, 5, 10, 10, 10, 5, 5, 5, 5, 5, 10, 20, 5, 5, 5, 5, 5, 20, 10, 10, 10, 5, 10, 20, 20, 20, 5, 5, 10, 10, 10, 20, 5, 20, 5, 20, 5, 5, 5, 5, 20, 5, 20, 5, 5, 20, 5, 10, 5, 20, 5, 5, 5, 5, 5, 5, 10, 5, 5, 10, 5, 20, 5, 5, 5, 20, 20, 10, 5, 5, 20, 5, 10, 5, 10, 5, 5, 5, 5, 20, 5, 20, 5, 10, 20, 10, 10, 5, 5, 5, 20, 5, 5, 5, 5, 5, 10, 5, 10, 5, 5, 5, 20, 20, 5, 5, 10, 10, 10, 5, 5, 5, 5, 10, 5, 5, 5, 5, 20, 20, 10, 5, 5, 5, 5, 5, 5, 20, 20, 5, 10, 5, 10, 10, 10, 10, 5, 10, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 10, 5, 5, 5, 5, 5, 20, 20, 10, 20, 10, 5, 5, 5, 20, 5, 5, 5, 20, 10, 5, 10, 5, 20, 5, 10, 5, 5, 5, 5, 5, 5, 20, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 10, 20, 10, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 10, 10, 5, 20, 5, 5, 5, 10, 5, 5, 10, 20, 10, 5, 20, 5, 20, 5, 5, 20, 10, 5, 20, 10, 5, 5, 5, 5, 10, 20, 5, 20, 20, 5, 5, 20, 10, 5, 5, 5, 10, 5, 10, 20, 5, 5, 5, 5, 5, 10, 10, 10, 10, 5, 10, 20, 5, 5, 20, 10, 10, 10, 20, 20, 20, 5, 20, 10, 5, 5, 10, 5, 5, 10, 5, 5, 5, 10, 5, 10, 5, 20, 20, 20, 5, 10, 10, 5, 5, 10, 5, 5, 5, 10, 20, 5, 10, 10, 10, 10, 20, 20, 5, 5, 20, 5, 5, 5, 5, 10, 5, 10, 5, 10, 5, 10, 10, 20, 5, 10, 5, 10, 20, 5, 5, 5, 10, 5, 10, 5, 5, 10, 5, 5, 5, 5, 10, 5, 20, 5, 5, 5, 5, 10, 10, 5, 20, 5, 5, 5, 10, 5, 10, 20, 5, 5, 20, 10, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 10, 20, 5, 20, 5, 10, 5, 5, 10, 5, 10, 5, 5, 5, 20, 10, 5, 5, 5, 5, 20, 10, 5, 5, 5, 5, 20, 5, 5, 5, 10, 5, 10, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 10, 5, 20, 10, 20, 5, 5, 20, 10, 5, 5, 10, 10, 10, 5, 10, 20, 5, 5, 5, 5, 10, 5, 20, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 20, 5, 20, 20, 5, 5, 10, 5, 5, 20, 10, 10, 5, 10, 20, 5, 10, 5, 5, 5, 10, 5, 5, 20, 5, 20, 20, 5, 5, 5, 20, 5, 5, 20, 5, 5, 5, 20, 5, 5, 5, 10, 5, 10, 10, 20, 20, 5, 5, 20, 5, 10, 20, 10, 5, 5, 5, 5, 10, 5, 5, 20, 5, 5, 5, 5, 5, 5, 10, 20, 20, 5, 5, 10, 20, 5, 5, 5, 10, 5, 5, 5, 20, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 20, 20, 5, 5, 5, 5, 5, 5, 20, 5, 5, 20, 10, 5, 10, 5, 5, 20, 10, 10, 5, 20, 20, 5, 5, 5, 5, 5, 5, 5, 20, 10, 5, 5, 5, 20, 10, 10, 5, 5, 5, 10, 10, 5, 10, 5, 5, 5, 20, 10, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 20, 10, 5, 5, 20, 20, 5, 20, 5, 10, 5, 10, 20, 10, 5, 10, 5, 5, 20, 5, 5, 5, 5, 20, 5, 10, 5, 5, 5, 20, 20, 5, 10, 10, 5, 10, 10, 20, 10, 5, 5, 10, 20, 5, 5, 10, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 5, 20, 10, 10, 5, 5, 5, 20, 5, 5, 20, 10, 5, 5, 5, 10, 5, 10, 5, 5, 5, 10, 5, 5, 20, 5, 10, 5, 10, 10, 5, 5, 10, 20, 20, 5, 5, 10, 10, 10, 5, 5, 5, 10, 5, 5, 5, 10, 5, 10, 5, 10, 20, 5, 20, 10, 20, 10, 5, 5, 5, 20, 5, 5, 20, 5, 5, 5, 5, 5, 10, 10, 5, 5, 5, 5, 10, 5, 5, 20, 20, 5, 10, 5, 5, 5, 5, 20, 5, 5, 20, 5, 10, 20, 20, 20, 5, 5, 20, 5, 10, 5, 5, 5, 5, 20, 10, 20, 10, 5, 5, 5, 10, 20, 5, 20, 5, 20, 10, 20, 20, 5, 10, 10, 5, 20, 5, 20, 5, 20, 5, 20, 5, 20, 20, 5, 5, 5, 5, 20, 20, 5, 5, 10, 5, 5, 20, 5, 5, 5, 10, 10, 5, 5, 5, 5, 5, 20, 20, 5, 10, 10, 10, 20, 5, 20, 20, 5, 5, 20, 20, 5, 5, 5, 10, 5, 20, 20, 10, 5, 5, 10, 5, 10, 10, 5, 5, 10, 5, 5, 5, 20, 5, 10, 5, 20, 5, 20, 5, 5, 20, 20, 10, 20, 5, 5, 5, 5, 5, 5, 20, 10, 5, 5, 5, 5, 5, 10, 20, 5, 5, 20, 5, 5, 5, 5, 5, 5, 20, 5, 5, 10, 10, 5, 5, 20, 20, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 10, 10, 5, 10, 10, 5, 5, 10, 10, 10, 20, 10, 10, 5, 5, 5, 10, 5, 5, 10, 10, 5, 10, 5, 5, 5, 10, 20, 5, 20, 20, 10, 5, 10, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 20, 5, 5, 20, 5, 5, 10, 20, 20, 20, 5, 5, 5, 5, 10, 10, 5, 5, 5, 5, 5, 20, 20, 5, 20, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 20, 5, 10, 20, 5, 5, 10, 5, 5, 5, 20, 20, 5, 10, 5, 10, 10, 5, 5, 10, 5, 5, 20, 20, 20, 20, 5, 5, 5, 10, 20, 10, 10, 10, 5, 10, 5, 5, 20, 5, 10, 5, 10, 5, 5, 20, 5, 5, 20, 5, 20, 5, 20, 20, 5, 20, 5, 5, 20, 5, 5, 5, 20, 5, 10, 5, 5, 20, 5, 20, 5, 20, 20, 20, 5, 5, 20, 5, 5, 10, 10, 5, 5, 5, 5, 10, 10, 5, 20, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 10, 5, 20, 5, 5, 20, 5, 10, 20, 5, 5, 10, 5, 5, 5, 5, 10, 5, 20, 5, 20, 5, 5, 5, 10, 10, 5, 5, 10, 5, 20, 20, 10, 5, 10, 5, 5, 5, 10, 5, 10, 5, 5, 10, 10, 5, 5, 20, 10, 5, 5, 5, 20, 5, 5, 20, 5, 5, 5, 20, 10, 10, 20, 10, 5, 10, 5, 5, 10, 5, 10, 10, 5, 5, 20, 10, 5, 10, 5, 5, 20, 20, 5, 20, 5, 5, 5, 5, 5, 20, 10, 5, 5, 5, 10, 10, 10, 5, 5, 20, 5, 5, 10, 5, 5, 5, 20, 5, 5, 5, 10, 5, 20, 10, 5, 5, 10, 5, 5, 10, 5, 5, 10, 5, 20, 5, 5, 5, 5, 20, 5, 20, 20, 5, 5, 5, 5, 5, 10, 10, 5, 5, 5, 20, 5, 20, 10, 5, 20, 5, 5, 5, 5, 5, 5, 10, 10, 5, 10, 10, 5, 5, 20, 5, 5, 20, 20, 10, 20, 5, 10, 5, 20, 10, 10, 10, 5, 5, 10, 5, 10, 5, 20, 5, 10, 5, 5, 10, 5, 5, 5, 10, 5, 5, 10, 10, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 10, 5, 5, 5, 5, 5, 5, 5, 20, 10, 5, 5, 10, 5, 5, 10, 10, 20, 20, 5, 10, 10, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 10, 20, 20, 5, 5, 20, 20, 5, 10, 20, 10, 5, 5, 10, 20, 10, 5, 20, 10, 10, 5, 5, 5, 5, 10, 5, 5, 5, 5, 20, 5, 5, 10, 10, 5, 20, 20, 5, 20, 10, 5, 20, 5, 5, 5, 20, 5, 5, 20, 20, 20, 5, 20, 20, 5, 10, 5, 10, 5, 5, 5, 10, 5, 20, 10, 5, 5, 5, 5, 20, 10, 5, 10, 10, 5, 5, 5, 5, 5, 5, 20, 20, 10, 10, 5, 10, 20, 20, 5, 20, 5, 20, 10, 5, 5, 5, 5, 10, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 5, 5, 5, 20, 10, 10, 20, 5, 5, 20, 10, 10, 5, 10, 5, 10, 5, 5, 20, 5, 5, 5, 5, 5, 10, 5, 10, 5, 5, 5, 5, 5, 20, 5, 10, 10, 10, 5, 5, 5, 10, 5, 20, 5, 5, 5, 10, 5, 5, 5, 5, 10, 5, 10, 5, 5, 10, 20, 5, 5, 10, 10, 20, 20, 20, 20, 10, 20, 5, 5, 5, 5, 5, 5, 5, 20, 5, 10, 5, 10, 20, 5, 20, 5, 5, 5, 5, 5, 20, 5, 5, 5, 20, 5, 20, 20, 20, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 20, 10, 5, 10, 10, 5, 5, 20, 20, 10, 5, 10, 10, 20, 10, 5, 5, 10, 10, 5, 10, 5, 5, 5, 10, 5, 5, 10, 10, 10, 5, 20, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 10, 5, 10, 10, 5, 5, 5, 20, 10, 5, 5, 10, 20, 20, 20, 5, 10, 10, 20, 20, 5, 20, 5, 5, 20, 5, 10, 5, 20, 20, 10, 5, 5, 10, 5, 5, 5, 10, 5, 5, 10, 10, 5, 20, 5, 10, 10, 5, 5, 10, 10, 5, 5, 5, 5, 10, 5, 10, 20, 5, 5, 10, 5, 20, 5, 20, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 20, 20, 20, 5, 20, 5, 20, 10, 5, 5, 5, 5, 5, 5, 20, 5, 20, 5, 5, 5, 5, 10, 5, 5, 5, 5, 20, 5, 5, 5, 10, 5, 5, 10, 20, 5, 20, 5, 10, 5, 20, 20, 10, 20, 5, 10, 5, 5, 10, 10, 20, 5, 20, 5, 20, 5, 20, 20, 20, 5, 20, 5, 20, 5, 5, 10, 20, 5, 10, 5, 5, 10, 5, 5, 20, 10, 10, 10, 5, 5, 5, 5, 5, 10, 5, 10, 10, 20, 5, 5, 5, 5, 5, 10, 10, 10, 5, 10, 20, 20, 5, 5, 20, 20, 5, 20, 5, 5, 5, 10, 10, 5, 5, 5, 10, 20, 5, 10, 5, 5, 10, 5, 10, 5, 10, 10, 5, 20, 5, 5, 10, 5, 5, 10, 10, 20, 5, 20, 20, 20, 5, 5, 10, 20, 5, 20, 5, 20, 5, 5, 10, 5, 20, 10, 20, 5, 10, 5, 10, 10, 5, 20, 5, 20, 5, 5, 5, 10, 5, 5, 20, 5, 5, 5, 5, 5, 10, 10, 20, 5, 20, 20, 5, 5, 10, 5, 5, 20, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 20, 10, 20, 5, 5, 5, 10, 5, 20, 5, 10, 5, 5, 10, 5, 5, 5, 10, 10, 5, 5, 10, 5, 5, 20, 20, 20, 5, 5, 20, 20, 10, 5, 5, 10, 5, 5, 5, 20, 5, 5, 20, 5, 20, 5, 20, 5, 5, 5, 10, 5, 20, 20, 5, 10, 10, 5, 5, 5, 5, 5, 5, 10, 20, 5, 5, 5, 20, 5, 5, 5, 20, 5, 5, 5, 5, 10, 20, 10, 5, 5, 5, 10, 5, 20, 20, 10, 10, 5, 10, 5, 20, 20, 5, 5, 5, 10, 5, 20, 5, 5, 20, 5, 5, 10, 20, 5, 20, 20, 5, 5, 20, 20, 20, 20, 5, 5, 10, 5, 20, 5, 5, 5, 5, 10, 5, 5, 5, 5, 10, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 10, 20, 20, 20, 5, 5, 10, 5, 5, 10, 5, 5, 20, 5, 10, 5, 10, 10, 5, 5, 20, 5, 10, 20, 5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 10, 5, 5, 10, 10, 5, 5, 20, 5, 10, 5, 20, 5, 20, 20, 5, 20, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 10, 20, 10, 5, 20, 20, 10, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 10, 20, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 20, 10, 5, 5, 10, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 20, 20, 5, 5, 20, 20, 10, 5, 10, 5, 10, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5, 10, 10, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 20, 5, 20, 5, 5, 5, 20, 5, 20, 5, 5, 5, 5, 5, 20, 5, 5, 20, 5, 10, 5, 5, 20, 5, 5, 5, 5, 5, 10, 10, 10, 5, 5, 5, 10, 10, 5, 20, 5, 10, 5, 5, 5, 5, 20, 20, 20, 5, 5, 5, 20, 20, 10, 5, 5, 5, 10, 5, 5, 5, 5, 5, 20, 5, 5, 20, 5, 10, 20, 5, 10, 5, 10, 10, 10, 10, 5, 10, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 20, 5, 5, 20, 5, 5, 5, 20, 5, 5, 5, 5, 5, 10, 10, 5, 5, 5, 5, 5, 20, 10, 10, 5, 10, 10, 20, 10, 5, 20, 10, 10, 5, 5, 20, 5, 5, 5, 5, 20, 5, 5, 5, 5, 10, 20, 5, 20, 20, 5, 20, 5, 10, 10, 5, 5, 10, 20, 5, 5, 5, 5, 20, 20, 5, 5, 5, 5, 10, 5, 5, 10, 5, 5, 5, 5, 5, 10, 5, 5, 10, 5, 10, 10, 10, 10, 5, 20, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 20, 10, 10, 5, 5, 5, 5, 5, 20, 5, 5, 5, 20, 5, 10, 10, 5, 5, 5, 5, 5, 20, 5, 20, 5, 5, 5, 5, 5, 10, 5, 10, 5, 5, 5, 5, 5, 10, 5, 5, 10, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 10, 20, 20, 5, 20, 20, 10, 10, 5, 5, 5, 5, 10, 10, 5, 10, 5, 10, 10, 5, 20, 5, 5, 20, 5, 5, 10, 5, 20, 5, 5, 5, 5, 20, 5, 5, 20, 5, 10, 5, 5, 20, 10, 10, 20, 10, 20, 20, 10, 20, 5, 10, 5, 20, 10, 5, 10, 5, 5, 10, 5, 5, 20, 5, 5, 10, 10, 10, 5, 5, 20, 5, 5, 5, 20, 5, 10, 10, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 10, 5, 5, 5, 20, 5, 20, 10, 10, 20, 20, 5, 20, 20, 10, 20, 5, 5, 5, 20, 5, 5, 5, 5, 20, 5, 10, 5, 10, 5, 5, 5, 20, 20, 5, 5, 10, 20, 10, 5, 5, 5, 10, 10, 5, 10, 20, 10, 10, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 20, 5, 20, 20, 5, 5, 5, 20, 20, 5, 5, 20, 5, 10, 20, 20, 5, 20, 10, 10, 5, 5, 5, 10, 5, 10, 5, 5, 5, 20, 5, 20, 10, 5, 5, 5, 5, 5, 5, 5, 10, 10, 10, 5, 10, 5, 5, 5, 20, 20, 20, 5, 5, 5, 5, 5, 20, 5, 20, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 10, 10, 5, 10, 20, 5, 10, 20, 10, 5, 5, 5, 5, 20, 5, 10, 5, 20, 5, 5, 5, 20, 5, 5, 10, 5, 5, 10, 5, 10, 10, 5, 5, 5, 5, 20, 5, 20, 20, 10, 5, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 5, 10, 10, 5, 5, 5, 20, 5, 5, 5, 20, 10, 20, 5, 5, 20, 10, 5, 5, 5, 20, 10, 5, 20, 5, 5, 5, 20, 5, 5, 5, 20, 5, 20, 5, 10, 5, 5, 10, 10, 10, 10, 5, 5, 5, 20, 20, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 20, 20, 5, 5, 20, 5, 5, 5, 5, 20, 5, 5, 10, 10, 5, 5, 5, 5, 10, 5, 5, 20, 5, 5, 5, 20, 5, 5, 5, 5, 20, 10, 5, 5, 5, 5, 10, 5, 5, 10, 20, 5, 5, 20, 5, 10, 5, 5, 20, 10, 5, 5, 20, 10, 20, 20, 5, 5, 20, 10, 10, 20, 5, 20, 5, 10, 10, 10, 5, 10, 5, 10, 5, 5, 5, 10, 5, 5, 10, 5, 5, 20, 20, 10, 10, 5, 5, 10, 20, 5, 5, 5, 10, 5, 10, 5, 5, 5, 5, 5, 5, 20, 5, 5, 10, 5, 5, 5, 5, 20, 5, 5, 5, 20, 5, 20, 5, 5, 5, 10, 5, 10, 20, 5, 20, 5, 10, 20, 10, 5, 5, 20, 10, 5, 10, 5, 10, 20, 20, 5, 20, 5, 5, 10, 5, 10, 20, 10, 20, 10, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 5, 10, 10, 5, 20, 5, 5, 5, 20, 10, 5, 10, 5, 5, 5, 5, 5, 10, 10, 5, 5, 10, 20, 20, 20, 5, 10, 5, 5, 5, 20, 5, 5, 5, 20, 5, 5, 5, 5, 10, 5, 5, 5, 5, 10, 20, 5, 20, 5, 5, 5, 5, 20, 20, 20, 5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 5, 10, 20, 5, 5, 5, 20, 5, 5, 10, 20, 20, 10, 5, 5, 5, 5, 20, 5, 5, 20, 20, 5, 5, 10, 20, 5, 5, 20, 5, 5, 10, 5, 10, 5, 10, 20, 10, 5, 5, 10, 5, 5, 5, 5, 5, 5, 20, 20, 20, 5, 5, 5, 5, 10, 20, 5, 10, 20, 5, 5, 5, 10, 20, 5, 10, 5, 5, 5, 5, 5, 10, 5, 5, 5, 20, 5, 5, 20, 5, 20, 5, 20, 10, 10, 5, 5, 20, 5, 5, 5, 5, 20, 20, 5, 5, 20, 20, 5, 5, 5, 10, 20, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 10, 5, 5, 5, 5, 10, 5, 5, 5, 20, 5, 5, 10, 5, 5, 5, 5, 20, 10, 10, 5, 10, 10, 5, 20, 5, 5, 5, 5, 10, 5, 20, 20, 5, 5, 10, 10, 20, 20, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 20, 5, 5, 10, 10, 20, 5, 20, 5, 20, 10, 5, 20, 10, 5, 10, 5, 5, 5, 5, 5, 10, 5, 5, 20, 5, 20, 5, 5, 20, 10, 10, 5, 10, 10, 5, 10, 5, 5, 5, 10, 10, 5, 5, 20, 10, 10, 10, 5, 5, 20, 5, 5, 20, 5, 5, 5, 20, 5, 5, 5, 10, 10, 5, 5, 5, 20, 20, 20, 5, 20, 5, 20, 5, 10, 5, 5, 5, 20, 5, 20, 10, 20, 5, 10, 5, 5, 20, 5, 5, 20, 10, 5, 5, 5, 5, 10, 10, 5, 5, 20, 20, 5, 10, 20, 5, 10, 5, 20, 5, 5, 10, 10, 5, 5, 5, 20, 5, 5, 5, 5, 10, 20, 5, 10, 5, 20, 5, 5, 5, 10, 5, 5, 5, 5, 20, 20, 5, 5, 10, 20, 5, 5, 5, 20, 5, 5, 10, 5, 10, 20, 10, 20, 10, 20, 5, 5, 20, 5, 5, 5, 5, 20, 10, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 20, 5, 20, 10, 5, 10, 10, 20, 10, 20, 5, 10, 10, 20, 5, 5, 5, 5, 5, 5, 5, 20, 20, 5, 20, 20, 20, 5, 5, 5, 5, 10, 5, 5, 5, 5, 20, 5, 5, 20, 10, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 10, 10, 5, 20, 10, 20, 10, 20, 5, 5, 5, 20, 5, 5, 5, 20, 10, 20, 5, 5, 5, 20, 5, 5, 5, 5, 5, 20, 5, 5, 10, 5, 5, 5, 5, 5, 20, 20, 5, 5, 5, 5, 5, 5, 5, 5, 20, 10, 5, 5, 20, 5, 5, 20, 20, 20, 5, 10, 10, 5, 5, 5, 5, 10, 20, 10, 10, 5, 5, 5, 5, 5, 20, 10, 5, 20, 10, 5, 10, 5, 5, 20, 5, 5, 5, 10, 5, 5, 5, 20, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 5, 10, 10, 5, 5, 5, 20, 5, 5, 10, 5, 5, 5, 20, 20, 20, 20, 5, 5, 5, 5, 10, 5, 20, 5, 5, 20, 20, 5, 5, 20, 10, 10, 5, 5, 5, 20, 5, 5, 5, 5, 5, 10, 5, 20, 20, 5, 5, 5, 10, 20, 5, 20, 5, 20, 5, 5, 5, 5, 5, 5, 5, 20, 20, 20, 5, 5, 20, 20, 10, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 20, 10, 20, 5, 5, 5, 5, 5, 5, 10, 5, 5, 10, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 10, 5, 5, 10, 10, 5, 10, 10, 5, 5, 20, 20, 5, 5, 5, 5, 5, 10, 5, 20, 10, 5, 5, 5, 5, 5, 5, 5, 5, 5, 10, 20, 20, 5, 20, 5, 20, 20, 20, 10, 20, 10, 5, 5, 10, 10, 20, 5, 5, 5, 5, 5, 10, 5, 10, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 20, 20, 20, 10, 5, 5, 10, 5, 5, 10, 5, 5, 20, 10, 20, 5, 5, 20, 20, 10, 20, 20, 5, 5, 5, 5, 5, 5, 5, 10, 5, 10, 10, 5, 10, 5, 20, 20, 10, 10, 5, 5, 20, 10, 5, 10, 10, 5, 5, 5, 5, 10, 5, 5, 20, 5, 5, 5, 5, 10, 5, 10, 5, 10, 5, 5, 5, 5, 10, 5, 5, 20, 5, 5, 5, 5, 5, 10, 20, 5, 10, 5, 10, 20, 10, 20, 20, 5, 5, 5, 5, 5, 20, 5, 10, 10, 5, 20, 5, 20, 10, 20, 10, 5, 10, 5, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 20, 5, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 10, 5, 10, 10, 20, 5, 20, 5, 5, 20, 5, 20, 5, 5, 5, 20, 10, 10, 5, 5, 5, 5, 5, 20, 10, 10, 5, 5, 5, 20, 10, 5, 10, 5, 5, 20, 10, 5, 5, 5, 5, 5, 5, 20, 5, 20, 5, 10, 10, 5, 5, 5, 5, 10, 5, 5, 5, 5, 10, 5, 10, 5, 10, 5, 10, 5, 5, 5, 5, 20, 5, 5, 20, 5, 5, 10, 5, 5, 5, 5, 10, 5, 5, 10, 5, 20, 5, 5, 5, 10, 10, 5, 20, 5, 5, 10, 5, 5, 10, 20, 5, 5, 5, 5, 10, 20, 20, 20, 5, 5, 10, 20, 5, 5, 5, 10, 5, 20, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 20, 5, 20, 5, 10, 5, 20, 20, 10, 10, 5, 5, 10, 5, 5, 5, 5, 5, 5, 10, 20, 10, 10, 5, 5, 5, 5, 5, 20, 5, 5, 5, 5, 5, 5, 5, 5, 10, 5, 20, 10, 20, 20, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 10, 20, 10, 20, 5, 10, 5, 5, 5, 5, 5, 5, 10, 10, 5, 5, 10, 20, 10, 5, 5, 5, 5, 20, 5, 5, 20, 5, 5, 10, 5, 5, 5, 5, 5, 5, 10, 20, 5, 20, 20, 10, 10, 10, 20, 10, 20, 5, 5, 10, 10, 5, 20, 5, 5, 10, 5, 5, 5, 5, 5, 5, 5, 10, 20, 10, 10, 5, 10, 20, 5, 5, 20, 5, 10, 5, 5, 10, 5, 20, 10, 10, 10, 5, 5, 5, 5, 5, 5, 5, 20, 5, 5, 20, 10, 5, 5, 20, 5, 5, 5, 5, 20, 20, 5, 10, 5, 20, 5, 5, 5, 5, 5, 5, 20, 10, 20, 20, 10, 5, 5, 5, 5, 5, 5, 20, 5, 10, 5, 5, 5, 5, 5, 10, 5, 5, 5, 20, 20, 5, 10, 20, 20, 20, 10, 20, 5, 5, 5, 5, 20, 5, 5, 5, 5, 20, 5, 5, 5, 5, 20, 5, 20, 5, 5, 5, 5, 5, 5, 10, 5, 5, 5, 5, 5, 10, 5, 10, 5, 10, 5, 5, 20, 10, 20, 10, 20, 5, 5, 5, 5, 5, 10, 20, 10, 20, 20, 10, 5, 5, 10, 5, 5, 10},
		result: false,
	})
	for _, test := range tests {
		if res := lemonadeChange(test.grid); res != test.result {
			t.Fatalf("error: %v  got: %v", test.result, res)
		}
	}
}

//lengthOfLongestSubstring
func TestLengthOfLongestSubstring(t *testing.T) {
	type lengthOfLongestSubstringTests struct {
		name   string
		grid   string
		result int
		lis    []int
	}
	var tests []lengthOfLongestSubstringTests
	tests = append(tests, lengthOfLongestSubstringTests{
		name:   "base test",
		grid:   "",
		result: 0,
	})
	tests = append(tests, lengthOfLongestSubstringTests{
		name:   "base test",
		grid:   "abcabcbb",
		result: 3,
	})
	tests = append(tests, lengthOfLongestSubstringTests{
		name:   "base test",
		grid:   "bbbbb",
		result: 1,
	})
	tests = append(tests, lengthOfLongestSubstringTests{
		name:   "base test",
		grid:   "pwwkew",
		result: 3,
	})
	tests = append(tests, lengthOfLongestSubstringTests{
		name:   "base test",
		grid:   "au",
		result: 2,
	})
	tests = append(tests, lengthOfLongestSubstringTests{
		name:   "base test",
		grid:   "aau",
		result: 2,
	})
	for _, test := range tests {
		if res := lengthOfLongestSubstring(test.grid); res != test.result {
			t.Fatalf("error: %v  got: %v", test.result, res)
		}
	}
}
