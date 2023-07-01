// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sort

import (
	"golang.org/x/exp/constraints"
)

// The time complexity of the merge sort algorithm is O(n*log(n)) in the average and worst cases, where n is the number of elements in the list.
// Merge sort is considered a stable sorting algorithm, meaning it preserves the relative order of elements with equal values.
func MergeSort[T constraints.Ordered](slice []T) []T {
	if len(slice) <= 1 {
		return slice
	}

	midIdx := len(slice) / 2
	left := MergeSort[T](slice[:midIdx])
	right := MergeSort[T](slice[midIdx:])
	return merge[T](left, right)
}

func merge[T constraints.Ordered](left, right []T) []T {
	if len(left) == 0 {
		return right
	}

	if len(right) == 0 {
		return left
	}

	res := make([]T, len(left)+len(right))
	var leftIdx, rightIdx int
	for leftIdx < len(left) || rightIdx < len(right) {
		if leftIdx >= len(left) {
			res = appendInPlace[T](right, res, rightIdx, leftIdx+rightIdx)
			break
		}

		if rightIdx >= len(right) {
			res = appendInPlace[T](left, res, leftIdx, leftIdx+rightIdx)
			break
		}

		if left[leftIdx] < right[rightIdx] {
			res[leftIdx+rightIdx] = left[leftIdx]
			leftIdx++
		} else {
			res[leftIdx+rightIdx] = right[rightIdx]
			rightIdx++
		}
	}
	return res
}

func appendInPlace[T constraints.Ordered](from, to []T, beginIndexFrom, beginIndexTo int) []T {
	for i, j := beginIndexTo, beginIndexFrom; j < len(from); i, j = i+1, j+1 {
		to[i] = from[j]
	}
	return to
}
