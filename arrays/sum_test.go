package main

import "testing"

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
	sum := []int{3, 9}

	got := SumAll(first_slice, second_slice)
	want := sum

	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
