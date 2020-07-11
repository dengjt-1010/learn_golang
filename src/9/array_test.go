package array_test

import (
	"testing"
)

func TestArrayInit(t *testing.T) {
	//var tarrat [3] int
	array2 := [...] int{1,2,3}
	t.Log(array2[2])
}

func TestArrayTravel(t *testing.T) {
	array3 := [...] int{1,3,4,5}
	for i := 0; i < len(array3); i++ {
		t.Log(array3[i])
	}

	for index,e:=range array3{
		t.Log(index,e)
	}
}

// 测试slice
func TestArraySlice(t *testing.T)  {
	array := [...] int{9,8,7,6,5,4,3,2,1}

	t.Log(array[:3])
	t.Log(array[6:])
	t.Log(array[2:3])
}

func TestArraySlice1(t *testing.T) {
	var slice1 []int

	t.Log(slice1)

	slice1 = append(slice1,1)
	slice1 = append(slice1,2)
	slice1 = append(slice1,3)

	t.Log(len(slice1), slice1[2])

	slice2 :=[]int{1,2,3,4}
	slice2 = append(slice2, 5)
	t.Log(len(slice2),cap(slice2),slice2[2])

	slice3 := make([]int,3,5)
	slice3 = append(slice3,1)
	slice3 = append(slice3,2)
	slice3 = append(slice3,3)
	t.Log("--------")
	t.Log(slice3[2],slice3[0])
}

func TestSlieceShareMemory(t *testing.T) {
	year:=[]int {1,2,3,4,5,6,7,8,9,10,11,12}

	Q2:= year[2:5]

	t.Log(Q2, len(Q2),cap(Q2))

	summer:=year[6:9]

	t.Log(summer,len(summer),cap(summer))

	summer[1] =100

	t.Log(summer,len(summer),cap(summer))
	t.Log(year)
	t.Log(year[7:8])
}

func TestSliceShareMemory2(t *testing.T) {
	year:=[]int {1,2,3,4,5,6,7,8,9,10,11,12}

	Q2:= year[2:5]
	Q2=append(Q2, 999)
	t.Log(year,len(year),cap(year))
	t.Log(Q2,len(Q2),cap(Q2))
	Q2[0]=1
	t.Log(year[2])

	//for i:=1;i<20;i++ {
	//	Q2 = append(Q2, i)
	//	//t.Log(Q2,len(Q2),cap(Q2),len(year),cap(year),)
	//}
}