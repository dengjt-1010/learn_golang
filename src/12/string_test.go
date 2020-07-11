package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s:="sbc"

	t.Log(len(s))//byte长度

	s="\xE4\xB8\xA5"
	t.Log(s)

	//s[2]='3'//不能复制，

	s="中华人民共和国"

	for _,c := range s{
		t.Logf("%[1]c %[1]d", c)
	}
}

func TestStringFn(t *testing.T) {

	s:="a,b,c"

	parts:=strings.Split(s,",")
	t.Log(parts)

	for index,value := range parts{
		t.Log(index,value)
	}

	t.Log(strings.Join(parts,"-"))
}

func TestStringConv(t *testing.T) {

	s:=strconv.Itoa(10)
	t.Log("this is a string",s)

	if i,error:=strconv.Atoi(s);error==nil{
		t.Log(i+10)
	}
}


