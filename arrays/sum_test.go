package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size(slice)", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got '%d' want '%d' given '%v'", got, want, numbers)
		}
	})

}

func TestSumAll(t *testing.T) {
	first_slice := []int{1, 2}
	second_slice := []int{0, 9}

	got := SumAll(first_slice, second_slice)
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	firstSlice := []int{1, 2}
	secondSlice := []int{0, 9}

	got := SumAllTails(firstSlice, secondSlice)
	want := []int{2, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
