package images

import (
	"errors"
)

var (
	ErrInvalidTransformationType = errors.New("invalid transformation type")
	ErrInvalidOrientation        = errors.New("invalid orientation")
	ErrParsingImage              = errors.New("error parsing image")
	ErrDecodeImage               = errors.New("error decoding image")
	ErrUnsupportedImageFormat    = errors.New("unsupported image format")
)
