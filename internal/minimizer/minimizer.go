package minimizer

import (
	"image"

	_ "image/jpeg"

	_ "image/gif"
	"image/png"
	"io"

	_ "golang.org/x/image/tiff"

	"golang.org/x/image/draw"
)

var errFormat error = image.ErrFormat

func (im *Imager) ErrFormat() error { return errFormat }

type Imager struct{}

func (im *Imager) Resize(in io.Reader, out io.Writer, lostgestLen int) error {
	img, _, err := image.Decode(in)
	if err != nil {
		return err
	}
	newX, newY := calculateSizes(img.Bounds().Dx(), img.Bounds().Dy(), lostgestLen)
	newImg := image.NewRGBA(image.Rect(0, 0, newX, newY))
	draw.BiLinear.Scale(newImg, newImg.Rect, img, img.Bounds(), draw.Over, nil)
	return png.Encode(out, newImg)
}

func longest(x, y int) int {
	return x - y
}

func calculateSizes(x, y, maxlen int) (int, int) {
	switch r := longest(x, y); {
	case r > 0:
		return maxlen, int(float64(maxlen) * float64(y) / float64(x))
	default:
		return int(float64(maxlen) * float64(x) / float64(y)), maxlen
	}
}
