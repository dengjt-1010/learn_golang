package try

import "testing"

const (
	Sencond = 1 + iota
	Third
	forth
)

func TestConstant(t *testing.T)  {
	t.Log(Sencond,Third,forth)
}

