package tableimage

import (
	"image"
)

const (
	rowSpace         = 30
	tablePadding     = 20
	letterPerPx      = 10
	separatorPadding = 10
)

type tableImage struct {
	width           int
	height          int
	columnSpace     int
	backgroundColor string
	img             *image.RGBA
}

func GenerateTableImage(ths []string, trs [][]string) {

	var totalRowNo = len(trs) + 1
	var totalColumnNo = len(ths)
	var columnSpace = getMaxAmountOfLetters(ths, trs) * letterPerPx
	var H = totalRowNo*rowSpace + rowSpace
	var W = totalColumnNo * int(columnSpace)

	ti := createContext(W, H, "#FFFFFF")
	ti.columnSpace = columnSpace
	ti.drawTH(ths)
	ti.drawTR(trs)

	ti.save("png", "./out.png")
}
