package tableimage

import (
	"fmt"
	"image/color"
	"strings"
)

func parseHexColor(x string) (r, g, b, a int) {

	x = strings.TrimPrefix(x, "#")
	a = 255
	if len(x) == 3 {
		format := "%1x%1x%1x"
		fmt.Sscanf(x, format, &r, &g, &b)
		r |= r << 4
		g |= g << 4
		b |= b << 4
	}
	if len(x) == 6 {
		format := "%02x%02x%02x"
		fmt.Sscanf(x, format, &r, &g, &b)
	}
	if len(x) == 8 {
		format := "%02x%02x%02x%02x"
		fmt.Sscanf(x, format, &r, &g, &b, &a)
	}
	return
}

func getColorByHex(hexColor string) color.RGBA {
	r, g, b, a := parseHexColor(hexColor)
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func getMaxAmountOfLetters(ths []string, trs [][]string) int {
	var maxColSpace = 0
	for _, th := range ths {
		if len(th) > maxColSpace {
			maxColSpace = len(th)
		}
	}

	for _, tr := range trs {
		for _, td := range tr {
			if len(td) > maxColSpace {
				maxColSpace = len(td)
			}
		}

	}

	return maxColSpace
}
