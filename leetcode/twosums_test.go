package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSums(t *testing.T) {
	for _, tc := range testCasesTwoSum {
		fmt.Println("tc: ", tc)
		t.Run(tc.description, func(t *testing.T) {
			actual := twoSum(tc.input, tc.target)
			if reflect.DeepEqual(actual, tc.expected) != true {
				t.Fatalf("twoSum(%v) = %v, want: %v", tc.input, actual, tc.expected)
			}
		})
	}
}
