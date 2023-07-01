// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sort

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
)

func TestQuickSortForInts(t *testing.T) {
	testCases := []struct {
		initial  []int
		expected []int
	}{
		{initial: []int{}, expected: []int{}},
		{initial: []int{1}, expected: []int{1}},
		{initial: []int{2, 1}, expected: []int{1, 2}},
		{initial: []int{1, 2, 3, 4, 5}, expected: []int{1, 2, 3, 4, 5}},
		{initial: []int{5, 4, 3, 2, 1}, expected: []int{1, 2, 3, 4, 5}},
		{initial: []int{1, 3, 4, 2, 5}, expected: []int{1, 2, 3, 4, 5}},
	}

	for i, tc := range testCases {
		assert.Exactlyf(t, tc.expected, QuickSort[int](tc.initial), msgWithIndex(i))
	}
}

func TestQuickSortForRandomInts(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	slice1, slice2 := copySlice[int](rand.Perm(10_000))
	sort.Ints(slice1)
	actual := QuickSort[int](slice2)
	assert.Exactly(t, slice1, actual)
}

func TestQuickSortForStrings(t *testing.T) {
	testCases := []struct {
		initial  []string
		expected []string
	}{
		{initial: []string{}, expected: []string{}},
		{initial: []string{"a"}, expected: []string{"a"}},
		{initial: []string{"a", "A"}, expected: []string{"A", "a"}},
		{initial: []string{"b", "a"}, expected: []string{"a", "b"}},
		{initial: []string{"a", "b", "c", "d", "e"}, expected: []string{"a", "b", "c", "d", "e"}},
		{initial: []string{"e", "d", "c", "b", "a"}, expected: []string{"a", "b", "c", "d", "e"}},
		{initial: []string{"a", "c", "d", "b", "e"}, expected: []string{"a", "b", "c", "d", "e"}},
	}

	for i, tc := range testCases {
		assert.Exactlyf(t, tc.expected, QuickSort[string](tc.initial), msgWithIndex(i))
	}
}

func TestQuickSortForRandomStrings(t *testing.T) {
	charset := "abcdefghijklmnopqrstuvwxyz"
	rand.Seed(time.Now().UnixNano())
	slice1 := make([]string, 10_000)
	for i := 0; i < 10_000; i++ {
		slice1[i] = string(charset[rand.Intn(len(charset))])
	}
	slice1, slice2 := copySlice[string](slice1)
	sort.Strings(slice1)
	actual := QuickSort[string](slice2)
	assert.Exactly(t, slice1, actual)
}

func copySlice[T constraints.Ordered](slice []T) ([]T, []T) {
	res := make([]T, len(slice))
	copy(res, slice)
	return slice, res
}

func msgWithIndex(i int) string {
	return fmt.Sprintf("test case index: %d", i)
}
