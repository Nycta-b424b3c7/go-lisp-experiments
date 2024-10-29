package main

import "testing"

func TestEq(t *testing.T) {
	{
		a := 1
		b := 1
		if !IsEq(a, b) {
			t.Fail()
		}
	}
	{
		a := 1
		b := 2
		if IsEq(a, b) {
			t.Fail()
		}
	}
}

func TestEqSymbol(t *testing.T) {
	{
		a := Symbol{"", "test"}
		b := Symbol{"", "test"}
		if !IsEq(a, b) {
			t.Fail()
		}
	}
	{
		a := Symbol{"", "test"}
		b := Symbol{"", "test_not"}
		if IsEq(a, b) {
			t.Fail()
		}
	}
}
