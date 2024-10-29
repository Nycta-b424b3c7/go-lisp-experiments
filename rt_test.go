package main

import "testing"

func TestTruthy(t *testing.T) {
	if Truthy(nil) {
		t.Fail()
	}

	if Truthy(false) {
		t.Fail()
	}

	if !Truthy(true) {
		t.Fail()
	}

	if !Truthy(1) {
		t.Fail()
	}

	if !Truthy("hello!") {
		t.Fail()
	}
}
