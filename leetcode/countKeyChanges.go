package leetcode

import (
	"strings"
)

func countKeyChanges(s string) int {
	lower := strings.ToLower(s)

	last := rune(lower[0])

	changes := 0
	for _, v := range lower {
		if v == last { //first time is the same
			continue
		} else {
			last = v
			changes = changes + 1
		}
	}
	return changes
}
