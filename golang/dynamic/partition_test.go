package dynamic

import (
	"testing"
)

func TestPartition(t *testing.T) {
	arr := []int{3, 1, 5, 9, 12}
	if !partition(arr) {
		t.FailNow()
	}
	arr = []int{3, 1, 1, 2, 2, 1}
	if !partition(arr) {
		t.FailNow()
	}
	arr = []int{3, 2, 1, 1, 2, 1}
	if !partition(arr) {
		t.FailNow()
	}
	arr = []int{3, 5}
	if partition(arr) {
		t.FailNow()
	}
}

func TestPartition2(t *testing.T) {
	arr := []int{3, 1, 5, 9, 12}
	if !partition2(arr) {
		t.FailNow()
	}
	arr = []int{3, 1, 1, 2, 2, 1}
	if !partition2(arr) {
		t.FailNow()
	}
	arr = []int{3, 2, 1, 1, 2, 1}
	if !partition2(arr) {
		t.FailNow()
	}
	arr = []int{3, 5}
	if partition2(arr) {
		t.FailNow()
	}
}
