package test

import (
	"bytes"
	"testing"
)

func AssertInt(t *testing.T, expected, received int) {
	if expected != received {
		t.Errorf("expected %d but received %d", expected, received)
	}
}

func AssertByteArray(t *testing.T, expected, received []byte) {
	if res := bytes.Compare(expected, received); res != 0 {
		t.Errorf("expected %s and received %s", expected, received)
	}
}
