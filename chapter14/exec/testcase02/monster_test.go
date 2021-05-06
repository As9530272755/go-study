package main

import (
	"testing"
)

func TestStore(t *testing.T) {
	var monster Monster
	if !monster.Store() {
		t.Fatalf("monster.Store() err 期望值 = %v , 实际值 = %v", true, monster.Store())
	}
	t.Logf("monster.Store()测试成功")
}

func TestReStore(t *testing.T) {
	var monster Monster
	if !monster.ReStore() {
		t.Fatalf("monster.ReStore() err 期望值 = %v , 实际值 = %v", true, monster.ReStore())
	}
	t.Logf("monster.Store()测试成功")
}
