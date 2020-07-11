package this_is_a__test

import "testing"

// TODO golang中，只有 a++ 或者 a--,符号不支持放在头部。

// todo 数组需要长度和内容都相同，并且位置一致
func TestOps(t *testing.T) {
	a:=[...] int{1,3,2}
	b:=[...] int{1,4,6}
	c:=[...] int{1,2,3}

	t.Log(a==b)
	t.Log(a==c)
}
