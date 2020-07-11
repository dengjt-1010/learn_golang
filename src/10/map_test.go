package testmap

import "testing"

func TestMapDefinition(t *testing.T) {
	m := map[string]int{"k1":1,"k2":2,"k3":3}
	t.Log(m["z"],m["k3"])//不存在 和 本身是0 区分不了？没有空指针异常

	t.Log(len(m))

	m2 :=map[string]int{}
	t.Log(m2)

	//这种方式可以区分map中是否有对应的key
	if v,ok:=m["k2"];ok{
		t.Log(v,ok)
	}else {
		t.Log("not exsist")
	}
}

func TestMapFor(t *testing.T) {
	m := map[string]int{"k1":1,"k2":2,"k3":3}
	for k,v := range m{
		t.Log(k,v)
	}

}
