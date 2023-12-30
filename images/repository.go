package images

import (
	"bytes"
	"image"
)

type Repository interface {
	EncodeImage(img image.Image, format string, buffer *bytes.Buffer) error
	Transform(img image.Image, transformation TransformationType) (image.Image, error)
	Transformation
}

type Transformation interface {
	Pixelate(img image.Image, pixelSize int) image.Image
	Stretch(img image.Image, size ImageSize) image.Image
	Jokerize(img image.Image) image.Image
	Chuuify(img image.Image) image.Image
}
