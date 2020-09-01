package _36

import "testing"

func BenchmarkSqurare(b *testing.B) {
	inputs := [...]int{1, 2, 3}
	expected := [...]int{1, 4, 9}
	b.StartTimer()
	for i := 0; i < len(inputs); i++ {
		ret := squre(inputs[i])
		if ret != expected[i] {
			b.Errorf("input is %d, expected is %d, the actual is %d", inputs[i], expected[i], ret)
		}
	}
	b.StopTimer()
}

func squre(input int) int {
	return input * input
}
