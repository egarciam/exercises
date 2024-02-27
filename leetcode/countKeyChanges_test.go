package leetcode

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	for _, tc := range testCases {
		fmt.Println("tc: ", tc)
		t.Run(tc.description, func(t *testing.T) {
			actual := countKeyChanges(tc.input)
			if actual != tc.expected {
				t.Fatalf("countKeyChanges(%q) = %dss, want: %d", tc.input, actual, tc.expected)
			}
		})
	}
}

// func countKeyChanges(s string) int {
// 	lower := strings.ToLower(s)

// 	last := rune(s[0])

// 	changes := 0
// 	for _, v := range lower {
// 		if last == v { //first time is the same
// 			continue
// 		} else {
// 			last = v
// 			changes = changes + 1
// 		}
// 	}
// 	if changes > 0 {
// 		changes = changes - 1
// 	}
// 	return changes
// }
