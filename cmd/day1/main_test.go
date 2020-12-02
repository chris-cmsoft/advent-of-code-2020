package main

import (
	"testing"
)

func TestSumInputs(t *testing.T) {
	if SumInputs([]int64{1,2,3}) != 6 {
		t.Errorf("SumInput not correctly calculating totals")
	}
}
