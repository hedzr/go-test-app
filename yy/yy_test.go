package yy_test

import (
	"github.com/hedzr/go-test-app/yy"
	"math/rand"
	"testing"
)

func TestOne(t *testing.T) {
	ret := yy.Factorial(3)
	if ret == 6 {
		t.Log(ret)
	} else {
		t.Fatal("bad")
	}
}

func TestFactorial(t *testing.T) {
	for i, tst := range []struct {
		input, expected int
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
	} {
		if ret := yy.Factorial(tst.input); ret != tst.expected {
			t.Fatalf("%3d. for Factorial(%d) expecting %v but got %v", i, tst.input, tst.expected, ret)
		}
	}
}

func testHelper(t *testing.T) {
	t.Helper()
	t.Skip()
	// ...
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rand.Int()
	}
}
