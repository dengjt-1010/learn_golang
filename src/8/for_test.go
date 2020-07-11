package testgo

import "testing"

func TestFor(t *testing.T) {
	a:=[...] int{1,3,2}
	for i:=0;i<len(a);i++{
		t.Log(a[i])
	}
}