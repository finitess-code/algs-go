// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// other tests are parameterized in sort_test.go

func TestAppendInPlace(t *testing.T) {
	to := make([]int, 6)
	to[0] = 1
	to[1] = 2
	to[2] = 3

	from := []int{3, 4, 5, 6}

	actual := appendInPlace[int](from, to, 1, 3)
	assert.Exactly(t, []int{1, 2, 3, 4, 5, 6}, actual)
}

func TestMerge(t *testing.T) {
	testCases := []struct {
		left     []int
		right    []int
		expected []int
	}{
		{left: []int{}, right: []int{}, expected: []int{}},
		{left: []int{1, 2}, right: []int{}, expected: []int{1, 2}},
		{left: []int{}, right: []int{2, 1}, expected: []int{2, 1}},
		{left: []int{1}, right: []int{1}, expected: []int{1, 1}},
		{left: []int{2, 1}, right: []int{4, 3}, expected: []int{2, 1, 4, 3}},
	}

	for i, tc := range testCases {
		assert.Exactlyf(t, tc.expected, merge[int](tc.left, tc.right), msgWithIndex(i))
	}
}
