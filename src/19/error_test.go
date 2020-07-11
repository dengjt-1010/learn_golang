package _9

import (
	"errors"
	"testing"
)

func fabnaci(n int) ([]int,error) {
	if n<0 || n>100{
		return nil,errors.New("n should be in [2,100]")
	}
	fblList := []int{1, 1}
	for i := 2; i < n; i++ {
		fblList = append(fblList, fblList[i-2]+fblList[i-1])
	}
	return fblList,nil
}

func TestGetFabnacci(t *testing.T) {
	if value,err:=fabnaci(1000);err==nil{
		t.Log("get fibanacci array", value)
	}else {
		t.Log("fail to get fibonacci",err)
	}
}
