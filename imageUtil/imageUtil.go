// (c)	2013-2014 Cergoo
// under terms of ISC license

// Image utils pkg
package imageUtil

import (
	"image"
	"image/draw"
)

type (
	IDistort interface {
		Init()
		Distort(src, dst *image.RGBA, dx, dy int)
	}
)

// Split image
func Split(img image.Image) (parts []*image.RGBA, hight int) {
	var (
		line  bool
		point [2]int
		rect  image.Rectangle
		a     uint32
	)

	hight = img.Bounds().Dy()
	width := img.Bounds().Dx()

	// find first and last points to lines
	for i := 0; i < width; i++ {
		_, _, _, a = img.At(i, 0).RGBA()
		if !line {
			if a > 0 {
				point[0] = i
				line = true
			}
		} else if a == 0 {
			point[1] = i
			line = false
			rect = image.Rect(0, 0, point[1]-point[0], hight)
			m := image.NewRGBA(rect)
			draw.Draw(m, rect, img, image.Point{point[0], 1}, draw.Src)
			parts = append(parts, m)
		}
	}

	return
}
