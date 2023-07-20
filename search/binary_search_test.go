// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package search

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearchForInts(t *testing.T) {
	testCases := []struct {
		initial  []int
		target   int
		expected bool
	}{
		{make([]int, 0), 5, false},
		{[]int{5}, 5, true},
		{[]int{5}, 2, false},
		{[]int{1, 3, 5, 7, 9}, 1, true},
		{[]int{1, 3, 5, 7, 9}, 9, true},
		{[]int{1, 3, 5, 7, 9}, 5, true},
		{[]int{1, 3, 5, 7, 9}, 2, false},
		{[]int{1, 3, 5, 7, 9}, 8, false},
	}

	for i, tc := range testCases {
		assert.Equalf(t, tc.expected, BinarySearch[int](tc.initial, tc.target), msgWithIndex(i))
	}
}

func TestBinarySearchForStrings(t *testing.T) {
	testCases := []struct {
		initial  []string
		target   string
		expected bool
	}{
		{make([]string, 0), "a", false},
		{[]string{"a"}, "a", true},
		{[]string{"a"}, "b", false},
		{[]string{"a", "c", "e", "g", "i"}, "a", true},
		{[]string{"a", "c", "e", "g", "i"}, "i", true},
		{[]string{"a", "c", "e", "g", "i"}, "e", true},
		{[]string{"a", "c", "e", "g", "i"}, "b", false},
		{[]string{"a", "c", "e", "g", "i"}, "h", false},
	}

	for i, tc := range testCases {
		assert.Equalf(t, tc.expected, BinarySearch[string](tc.initial, tc.target), msgWithIndex(i))
	}
}

func msgWithIndex(i int) string {
	return fmt.Sprintf("test case index: %d", i)
}
