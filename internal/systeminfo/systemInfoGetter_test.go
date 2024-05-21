package systeminfo

import "testing"

func TestCPU(t *testing.T) {
	_, err := CPU()
	if err != nil {
		t.Fatalf("gor error %v", err)
	}
}
