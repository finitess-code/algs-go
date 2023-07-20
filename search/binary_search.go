// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package search

import "golang.org/x/exp/constraints"

// Time complexity of O(log n), which makes it much faster than linear (O(n)) search for large data sets.
// Only works correctly on sorted collections.
func BinarySearch[T constraints.Ordered](slice []T, target T) bool {
	lowIdx := 0
	highIdx := len(slice) - 1

	for lowIdx <= highIdx {
		midIdx := (lowIdx + highIdx) / 2
		if slice[midIdx] == target {
			return true
		} else if slice[midIdx] < target {
			lowIdx = midIdx + 1
		} else {
			highIdx = midIdx - 1
		}
	}
	return false
}
