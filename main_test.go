package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Add(t *testing.T) {
	testcases := []struct {
		msg    string
		a      int
		b      int
		expect int
	}{
		{msg: "zero", a: 0, b: 0, expect: 0},
		{msg: "1+0", a: 1, b: 0, expect: 1},
		{msg: "1+1", a: 1, b: 1, expect: 2},
	}

	for _, tc := range testcases {
		t.Run(tc.msg, func(t *testing.T) {
			actual := Add(tc.a, tc.b)
			assert.Equal(t, tc.expect, actual)
		})
	}
}
