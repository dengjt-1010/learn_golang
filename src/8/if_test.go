package testgo

import (
	"runtime"
	"testing"
)

func TestIfMultiSec(t *testing.T)  {
	if a:=1==1; a{
		t.Log("if check is true")
	}
}

func TestGoSwitch(t *testing.T) {
	switch os:= runtime.GOOS;os {
	case "darwin":
		t.Log("Os X")
	case "Mac OS":
		t.Log("mac os")
	default:
		t.Log("oter sysh")
	}
}

func TestCaseToIf(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i<3 :
			t.Log("i less than 3")
		case i >= 3:
			t.Log("i equals or bigger than 3")
		}
	}
}

