// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sort

import (
	"golang.org/x/exp/constraints"
)

// Average time complexity O(n*log(n)), where n is the number of elements in the array.
// However, in the worst-case scenario, the time complexity can degrade to O(n^2) if the pivot selection is not optimal.
func QuickSort[T constraints.Ordered](slice []T) []T {
	if len(slice) <= 1 {
		return slice
	}

	pivotIdx := len(slice) / 2
	left := make([]T, 0)
	right := make([]T, 0)
	for i := range slice {
		if i == pivotIdx {
			continue
		}

		if slice[i] <= slice[pivotIdx] {
			left = append(left, slice[i])
		} else {
			right = append(right, slice[i])
		}
	}

	slicesToConcat := [][]T{QuickSort[T](left), {slice[pivotIdx]}, QuickSort[T](right)}
	return concat(slicesToConcat)
}

func concat[T constraints.Ordered](slices [][]T) []T {
	var totalLen int
	for _, s := range slices {
		totalLen += len(s)
	}

	res := make([]T, totalLen)

	var idx int
	for _, s := range slices {
		idx += copy(res[idx:], s)
	}

	return res
}
