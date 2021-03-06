package sliceutils

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDifference(t *testing.T) {
	cases := []struct {
		slice1, slice2 []string
		expectedSlice  []string
	}{
		{
			slice1:        []string{},
			slice2:        []string{},
			expectedSlice: []string{},
		},
		{
			slice1:        []string{"A"},
			slice2:        []string{},
			expectedSlice: []string{"A"},
		},
		{
			slice1:        []string{},
			slice2:        []string{"A"},
			expectedSlice: []string{},
		},
		{
			slice1:        []string{"A"},
			slice2:        []string{"B"},
			expectedSlice: []string{"A"},
		},
		{
			slice1:        []string{"A"},
			slice2:        []string{"A"},
			expectedSlice: nil,
		},
		{
			slice1:        []string{"A", "B", "C"},
			slice2:        []string{"B"},
			expectedSlice: []string{"A", "C"},
		},
		{
			slice1:        []string{"A", "B", "A", "C", "B"},
			slice2:        []string{"B"},
			expectedSlice: []string{"A", "C"},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%s - %s", strings.Join(c.slice1, " "), strings.Join(c.slice2, " ")), func(t *testing.T) {
			assert.Equal(t, c.expectedSlice, StringDifference(c.slice1, c.slice2))
		})
	}
}

func TestUnion(t *testing.T) {
	cases := []struct {
		slice1, slice2 []string
		expectedSlice  []string
	}{
		{
			slice1:        []string{},
			slice2:        []string{},
			expectedSlice: []string{},
		},
		{
			slice1:        []string{"A"},
			slice2:        []string{},
			expectedSlice: []string{"A"},
		},
		{
			slice1:        []string{},
			slice2:        []string{"A"},
			expectedSlice: []string{"A"},
		},
		{
			slice1:        []string{"A"},
			slice2:        []string{"B"},
			expectedSlice: []string{"A", "B"},
		},
		{
			slice1:        []string{"A"},
			slice2:        []string{"A"},
			expectedSlice: []string{"A"},
		},
		{
			slice1:        []string{"A", "A"},
			slice2:        []string{},
			expectedSlice: []string{"A"},
		},
		{
			slice1:        []string{"A", "A"},
			slice2:        []string{"B"},
			expectedSlice: []string{"A", "B"},
		},
		{
			slice1:        []string{"A", "B", "C"},
			slice2:        []string{"B"},
			expectedSlice: []string{"A", "B", "C"},
		},
		{
			slice1:        []string{"A", "B", "A", "C"},
			slice2:        []string{},
			expectedSlice: []string{"A", "B", "C"},
		},
		{
			slice1:        nil,
			slice2:        []string{"A", "B", "A", "C"},
			expectedSlice: []string{"A", "B", "C"},
		},
	}
	for _, c := range cases {
		t.Run(fmt.Sprintf("%s - %s", strings.Join(c.slice1, " "), strings.Join(c.slice2, " ")), func(t *testing.T) {
			// Run multiple times to ensure deterministic output.
			for i := 0; i < 10; i++ {
				assert.Equal(t, c.expectedSlice, StringUnion(c.slice1, c.slice2))
			}
		})
	}
}

func TestSelect(t *testing.T) {
	input := []string{"foo", "bar", "baz", "qux"}
	cases := []struct {
		indices  []int
		expected []string
		panics   bool
	}{
		{
			indices:  []int{1, 3},
			expected: []string{"bar", "qux"},
		},
		{
			indices:  []int{2, 0},
			expected: []string{"baz", "foo"},
		},
		{
			indices:  []int{},
			expected: nil,
		},
		{
			indices:  []int{0, 0, 1, 1, 2, 2, 3, 3},
			expected: []string{"foo", "foo", "bar", "bar", "baz", "baz", "qux", "qux"},
		},
		{
			indices: []int{0, -1},
			panics:  true,
		},
		{
			indices: []int{0, 4},
			panics:  true,
		},
	}

	for _, testCase := range cases {
		c := testCase
		t.Run(fmt.Sprintf("%v", c.indices), func(t *testing.T) {
			if c.panics {
				assert.Panics(t, func() {
					StringSelect(input, c.indices...)
				})
			} else {
				result := StringSelect(input, c.indices...)
				assert.Equal(t, c.expected, result)
			}
		})
	}
}
