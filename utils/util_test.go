package utils

import "testing"

func TestGetStartDateAndEndDate(t *testing.T) {
	a, b := GetStartDateAndEndDate()
	t.Logf("s=%v,e=%v,", a, b)
}

func TestRandN(t *testing.T) {
	t.Logf("%v",RandN(2))
}
