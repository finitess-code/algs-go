// Copyright 2023 Finitess.com. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// other tests are parameterized in sort_test.go

func TestConcat(t *testing.T) {
	slices := [][]int{
		{1, 2},
		{},
		{3},
		{4, 5, 6},
	}

	actual := concat[int](slices)
	assert.Exactly(t, []int{1, 2, 3, 4, 5, 6}, actual)
}
