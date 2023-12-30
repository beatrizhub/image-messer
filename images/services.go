package images

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

type Service struct {
}

func NewService() Repository {
	return &Service{}
}

func (s *Service) Transform(img image.Image, transformation TransformationType) (image.Image, error) {
	switch transformation {
	case stretched:
		img = s.Stretch(img, ImageSize{Width: img.Bounds().Dx() * 5, Height: img.Bounds().Dy()})
	case pixelated:
		img = s.Pixelate(img, 15)
	case joker:
		img = s.Jokerize(img)
	case chuu:
		img = s.Chuuify(img)
	default:
		return nil, ErrInvalidTransformationType
	}
	return img, nil
}

func (s *Service) Pixelate(img image.Image, pixelSize int) image.Image {
	return resize.Resize(uint(img.Bounds().Dx()/pixelSize), uint(img.Bounds().Dy()/pixelSize), img, resize.NearestNeighbor)
}

func (s *Service) Stretch(img image.Image, size ImageSize) image.Image {
	return resize.Resize(uint(size.Width), uint(size.Height), img, resize.Lanczos3)
}

func (s *Service) Jokerize(img image.Image) image.Image {
	jokerImg, err := loadImage(jokerPath)
	if err != nil {
		return nil
	}

	jokerizedImg, err := combineImages(img, jokerImg, "vertical")
	if err != nil {
		return nil
	}

	return jokerizedImg
}

func (s *Service) Chuuify(img image.Image) image.Image {
	chuuImg, err := loadImage(chuuPath)
	if err != nil {
		return nil
	}

	chuuifiedImg, err := combineImages(img, chuuImg, "vertical")
	if err != nil {
		return nil
	}

	return chuuifiedImg
}

func (s *Service) EncodeImage(img image.Image, format string, buffer *bytes.Buffer) error {
	switch format {
	case "png":
		png.Encode(buffer, img)
	case "jpeg", "jpg":
		jpeg.Encode(buffer, img, nil)
	default:
		return ErrUnsupportedImageFormat
	}

	return nil
}

func (s *Service) Save(path string, img image.Image) error {
	b := bytes.NewBuffer(nil)

	err := jpeg.Encode(b, img, nil)
	if err != nil {
		return err
	}

	newImg, err := os.Create(path)
	if err != nil {
		return err
	}
	defer newImg.Close()

	_, err = newImg.Write(b.Bytes())
	if err != nil {
		return err
	}

	return nil
}

func combineImages(originalImage image.Image, overlayImage image.Image, orientation string) (image.Image, error) {
	var combinedImage *image.RGBA

	switch orientation {
	case "vertical":
		combinedImage = combineVertically(originalImage, overlayImage)
	case "horizontal":
		combinedImage = combineHorizontally(originalImage, overlayImage)
	default:
		return nil, ErrInvalidOrientation
	}

	return combinedImage, nil
}

func combineVertically(originalImage image.Image, overlayImage image.Image) *image.RGBA {
	overlayImage = resize.Resize(uint(originalImage.Bounds().Dx()/2), uint(originalImage.Bounds().Dy()), overlayImage, resize.Lanczos3)
	combinedImage := image.NewRGBA(image.Rect(0, 0, originalImage.Bounds().Dx(), originalImage.Bounds().Dy()))

	for x := originalImage.Bounds().Dx() / 2; x < originalImage.Bounds().Dx(); x++ {
		for y := 0; y < originalImage.Bounds().Dy(); y++ {
			r, g, b, a := originalImage.At(x, y).RGBA()
			combinedImage.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}

	for x := 0; x < originalImage.Bounds().Dx()/2; x++ {
		for y := 0; y < originalImage.Bounds().Dy(); y++ {
			r, g, b, a := overlayImage.At(x, y).RGBA()
			combinedImage.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}

	return combinedImage
}

func combineHorizontally(originalImage image.Image, overlayImage image.Image) *image.RGBA {
	overlayImage = resize.Resize(uint(originalImage.Bounds().Dx()), uint(originalImage.Bounds().Dy()/2), overlayImage, resize.Lanczos3)
	combinedImage := image.NewRGBA(image.Rect(0, 0, originalImage.Bounds().Dx(), originalImage.Bounds().Dy()))

	for x := 0; x < originalImage.Bounds().Dx(); x++ {
		for y := originalImage.Bounds().Dy() / 2; y < originalImage.Bounds().Dy(); y++ {
			r, g, b, a := originalImage.At(x, y).RGBA()
			combinedImage.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}

	for x := 0; x < originalImage.Bounds().Dx(); x++ {
		for y := 0; y < originalImage.Bounds().Dy()/2; y++ {
			r, g, b, a := overlayImage.At(x, y).RGBA()
			combinedImage.Set(x, y, color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)})
		}
	}

	return combinedImage
}

func loadImage(path string) (image.Image, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}

	return img, nil
}
