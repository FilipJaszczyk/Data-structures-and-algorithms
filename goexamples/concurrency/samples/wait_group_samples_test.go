package samples

import "testing"

func TestWgBasicUsage(t *testing.T) {
	if res := WgBasicUsage(); res != 1 {
		t.Fatalf("expected %d", 1)
	}
}
