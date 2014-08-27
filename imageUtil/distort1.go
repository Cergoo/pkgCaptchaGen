// (c)	2013-2014 Cergoo
// under terms of ISC license

// Package imageUtil it's util
package imageUtil

import (
	"image"
	"math"
	"math/rand"
)

type (
	// TDistort1 it's IDistor interface implementation
	TDistort1 struct {
		liney, delta int
		part1        struct{ d1, d2 int }
		part2        struct{ d1, d2 int }
	}
)

// Init init random value distort
func (t *TDistort1) Init() {
	t.liney = 20 + rand.Intn(10)
	t.delta = 10

	t.part1.d1 = 2 + rand.Intn(4)
	t.part1.d2 = t.part1.d1*5 + rand.Intn(5)

	t.part2.d1 = 30
	t.part2.d2 = 75 + rand.Intn(5)
	if rand.Intn(2) == 0 {
		t.part2.d1 *= -1
		t.delta = 40
	}
}

// Distort distort img
func (t *TDistort1) Distort(src, dst *image.RGBA, dx, dy int) {
	var (
		dstx float64
		a    uint32
		x, y int
	)

	for y = 0; y < src.Bounds().Dy(); y++ {
		if y == t.liney {
			dx++
			continue
		}
		for x = 0; x < src.Bounds().Dx(); x++ {
			if _, _, _, a = src.At(x, y).RGBA(); a > 0 {
				dstx = float64(dx + x + t.delta)
				dstx += float64(t.part2.d1) * math.Sin(6.28*float64(y)/float64(t.part2.d2))
				dstx += float64(t.part1.d1) * math.Sin(6.28*float64(y)/float64(t.part1.d2))
				dst.Set(int(dstx), y, src.At(x, y))
			}
		}
	}
}
