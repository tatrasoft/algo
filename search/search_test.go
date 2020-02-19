package search

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumInList_NumIsInTheList(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	num := 5

	isInList := NumInList(list, num)
	assert.True(t, isInList)
}

func TestNumInList_NumIsInTheListWithSameNumbers(t *testing.T) {
	list := []int{5, 5, 5, 5, 5}
	num := 5

	isInList := NumInList(list, num)
	assert.True(t, isInList)
}

func TestNumInList_NumIsInTheListWithNegatives(t *testing.T) {
	list := []int{1, 2, -3, 4, 5}
	num := -3

	isInList := NumInList(list, num)
	assert.True(t, isInList)
}

func TestNumInList_NumIsNotInTheList(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	num := 10

	isInList := NumInList(list, num)
	assert.False(t, isInList)
}

func TestNumInList_NilSlice(t *testing.T) {
	num := 5

	isInList := NumInList(nil, num)
	assert.False(t, isInList)
}

func TestNumInList_EmptySlice(t *testing.T) {
	list := []int{}
	num := 5

	isInList := NumInList(list, num)
	assert.False(t, isInList)
}


