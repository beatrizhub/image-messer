package images

import (
	"image"
	"strings"
)

type ImageSize struct {
	Bound         image.Rectangle
	Width, Height int
}

type TransformationType string

const (
	stretched TransformationType = "stretched"
	pixelated TransformationType = "pixelated"
	joker     TransformationType = "joker"
	chuu      TransformationType = "chuu"

	jokerPath = "assets/transform/joker.png"
	chuuPath  = "assets/transform/chuu.png"
)

func NewTransformationType(tt string) TransformationType {
	switch strings.ToLower(tt) {
	case "stretched":
		return stretched
	case "pixelated":
		return pixelated
	case "joker":
		return joker
	case "chuu":
		return chuu
	default:
		return ""
	}
}
