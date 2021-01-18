package main

import (
	"reflect"
	"testing"
)

func TestExpenseReportSum(t *testing.T) {
	items := []int{1721, 979, 366, 299, 675, 1456}

	got := expenseReportSum(items, 2020)
	want := []int{299, 1721}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got: %v; want: %v", got, want)
	}
}
