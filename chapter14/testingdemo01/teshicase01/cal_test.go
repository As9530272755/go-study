package main

import (
	"testing"
)

func TestSub(t *testing.T) {
	res := Sub(10, 4)
	n := 6
	if res != n {
		t.Fatalf("测试错误 期望值 = %v 实际值 = %v", n, res)
	} else {
		t.Logf("测试正确 实际值 = %v", res)
	}
}

func TestAddUpper(t *testing.T) {
	res := AddUpper(10)

	if res != 55 {
		t.Fatalf("Addupper 测试错误，错误值 = %v", res)
	} else {
		t.Logf("AddUpper 测试正确 = %v", res)
	}
}
