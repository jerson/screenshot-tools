package utils

import (
	"errors"
	"image"
	"image/color"
	"log"
	"math"
	"os"
	"screenshot_tools/modules/config"

	"github.com/disintegration/imaging"
)

// ImageFixed ...
type ImageFixed struct {
	Image *image.NRGBA
	Point image.Point
}

// MergeImages ...
func MergeImages(paths []string, output string) error {

	if len(paths) < 1 {
		return errors.New("empty paths")
	}
	resizeHeight := config.Vars.Resize.Height
	x := 0
	y := 0
	width := 0
	height := resizeHeight

	var images []ImageFixed
	for i, path := range paths {
		img, err := decode(path)
		if err != nil {
			return err
		}

		src := imaging.Resize(img, 0, resizeHeight, imaging.Lanczos)
		images = append(images, ImageFixed{
			Point: image.Pt(x, y),
			Image: src,
		})
		x += src.Bounds().Max.X
		if x > width {
			width = x
		}
		if math.Mod(float64(i+1), float64(config.Vars.Resize.Columns)) == 0 {
			x = 0
			y += resizeHeight
			if i < len(paths)-1 {
				height += resizeHeight
			}
		}

	}

	target := imaging.New(width, height, color.NRGBA{0, 0, 0, 0})
	for _, imageFixed := range images {
		target = imaging.Paste(target, imageFixed.Image, imageFixed.Point)
	}
	err := imaging.Save(target, output)

	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	return nil
}

func decode(path string) (image.Image, error) {
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
