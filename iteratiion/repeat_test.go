package iteratiion

import (
	"fmt"
	"testing"
)

func TestIteration(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %s, repeated %s", expected, repeated)
	}
}

func ExampleRepeat() {
	repeatedChrs := Repeat("b", 5)
	fmt.Println(repeatedChrs)
	// Output: bbbbb
}

func BenchmarkRepat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
