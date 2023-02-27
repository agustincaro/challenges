package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1_emptySlice_returnFalse(t *testing.T) {

	var emptySlice = []int{}
	var k = 10

	expected := false
	actual := listAdds(emptySlice, k)

	assert.Equal(t, expected, actual)
}

func Test_slice_two_elem_return_true(t *testing.T) {
	slice := []int{10, 7}
	k := 17

	expected := true
	actual := listAdds(slice, k)

	assert.Equal(t, expected, actual)
}

func Test_slice_two_elem_retun_false(t *testing.T) {
	slice := []int{10, 7}
	k := 18

	expected := false
	actual := listAdds(slice, k)

	assert.Equal(t, expected, actual)
}

func Test_slice_multiple_elem_retun_true(t *testing.T) {
	slice := []int{10, 7, 12, 15, 4, 1}
	k := 5

	expected := true
	actual := listAdds(slice, k)

	assert.Equal(t, expected, actual)
}

func Test_slice_multiple_elem_return_false(t *testing.T) {
	slice := []int{10, 7, 12, 15, 4, 1}
	k := 40

	expected := false
	actual := listAdds(slice, k)

	assert.Equal(t, expected, actual)
}
