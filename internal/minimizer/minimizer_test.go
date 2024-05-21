package minimizer

import (
	"os"
	"testing"
)

func TestResize(t *testing.T) {
	src, err := os.Open("source.jpeg")
	if err != nil {
		t.Fatal(err)
	}
	dest, err := os.OpenFile("dest.jpeg", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		t.Fatal(err)
	}
	imager := &Imager{}
	err = imager.Resize(src, dest, 551)
	if err != nil {
		t.Fatal(err)
	}
}
func BenchmarkCalculateSizes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		calculateSizes(400, 200, 1200)
	}
}
