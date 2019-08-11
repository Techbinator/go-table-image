package tableimage

import (
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

func createContext(width int, height int, backgroundColor string) tableImage {
	ti := tableImage{
		width:           width,
		height:          height,
		backgroundColor: backgroundColor,
	}
	ti.setRgba()
	return ti
}

func (ti *tableImage) setRgba() {
	img := image.NewRGBA(image.Rect(0, 0, ti.width, ti.height))
	//set image background
	draw.Draw(img, img.Bounds(), &image.Uniform{getColorByHex(ti.backgroundColor)}, image.ZP, draw.Src)
	ti.img = img
}

func (ti *tableImage) addString(x, y int, label string, color string) {

	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}

	d := &font.Drawer{
		Dst:  ti.img,
		Src:  image.NewUniform(getColorByHex(color)),
		Face: basicfont.Face7x13,
		Dot:  point,
	}
	d.DrawString(label)
}

func (ti *tableImage) addLine(x1, y1, x2, y2 int, color string) {
	//Thx to https://github.com/StephaneBunel/bresenham
	var dx, dy, e, slope int
	col := getColorByHex(color)
	// Because drawing p1 -> p2 is equivalent to draw p2 -> p1,
	// I sort points in x-axis order to handle only half of possible cases.
	if x1 > x2 {
		x1, y1, x2, y2 = x2, y2, x1, y1
	}

	dx, dy = x2-x1, y2-y1
	// Because point is x-axis ordered, dx cannot be negative
	if dy < 0 {
		dy = -dy
	}

	switch {

	// Is line a point ?
	case x1 == x2 && y1 == y2:
		ti.img.Set(x1, y1, col)

	// Is line an horizontal ?
	case y1 == y2:
		for ; dx != 0; dx-- {
			ti.img.Set(x1, y1, col)
			x1++
		}
		ti.img.Set(x1, y1, col)

	// Is line a vertical ?
	case x1 == x2:
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		for ; dy != 0; dy-- {
			ti.img.Set(x1, y1, col)
			y1++
		}
		ti.img.Set(x1, y1, col)

	// Is line a diagonal ?
	case dx == dy:
		if y1 < y2 {
			for ; dx != 0; dx-- {
				ti.img.Set(x1, y1, col)
				x1++
				y1++
			}
		} else {
			for ; dx != 0; dx-- {
				ti.img.Set(x1, y1, col)
				x1++
				y1--
			}
		}
		ti.img.Set(x1, y1, col)

	// wider than high ?
	case dx > dy:
		if y1 < y2 {
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				ti.img.Set(x1, y1, col)
				x1++
				e -= dy
				if e < 0 {
					y1++
					e += slope
				}
			}
		} else {
			dy, e, slope = 2*dy, dx, 2*dx
			for ; dx != 0; dx-- {
				ti.img.Set(x1, y1, col)
				x1++
				e -= dy
				if e < 0 {
					y1--
					e += slope
				}
			}
		}
		ti.img.Set(x2, y2, col)

	// higher than wide.
	default:
		if y1 < y2 {
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				ti.img.Set(x1, y1, col)
				y1++
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		} else {
			dx, e, slope = 2*dx, dy, 2*dy
			for ; dy != 0; dy-- {
				ti.img.Set(x1, y1, col)
				y1--
				e -= dx
				if e < 0 {
					x1++
					e += slope
				}
			}
		}
		ti.img.Set(x2, y2, col)
	}
}

func (ti *tableImage) save(fileType string, filePath string) {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if fileType == "jpg" {
		if err := jpeg.Encode(f, ti.img, nil); err != nil {
			panic(err)
		}
	} else {
		if err := png.Encode(f, ti.img); err != nil {
			panic(err)
		}
	}
}
