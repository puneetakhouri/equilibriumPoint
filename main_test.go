package main

import "testing"

func TestFindEquilibrium(t *testing.T) {
	dataTable := []struct {
		arr      []int
		expected int
	}{
		{[]int{1}, 1},
		{[]int{1, 3, 5, 2, 2}, 3},
	}

	for _, data := range dataTable {
		if data.expected == FindEquilibrium(data.arr) {
			t.Log("SUCCESS")
		} else {
			t.Error("Error")
		}

	}
}
