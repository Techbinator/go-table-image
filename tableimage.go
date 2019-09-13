package tableimage

import (
	"image"
)

//FileType the image format pmng or jpg
type FileType string

//TD a table data container
type TD struct {
	Text  string
	Color string
}

//TR a table row
type TR struct {
	BorderColor string
	Tds         []TD
}

type tableImage struct {
	width           int
	height          int
	th              TR
	trs             []TR
	backgroundColor string
	fileType        FileType
	filePath        string
	img             *image.RGBA
}

const (
	rowSpace                  = 26
	tablePadding              = 20
	letterPerPx               = 10
	separatorPadding          = 10
	wrapWordsLen              = 20
	columnSpace               = wrapWordsLen * letterPerPx
	PNG              FileType = "png"
	JPEG             FileType = "jpg"
)

func Init(backgroundColor string, fileType FileType, filePath string) tableImage {
	ti := tableImage{
		backgroundColor: backgroundColor,
		fileType:        fileType,
		filePath:        filePath,
	}
	ti.setRgba()
	return ti
}

func (ti *tableImage) AddTH(th TR) {
	ti.th = th
}

func (ti *tableImage) AddTRs(trs []TR) {
	ti.trs = trs
}

func (ti *tableImage) Save() {
	ti.calculateHeight()
	ti.calculateWidth()

	ti.setRgba()

	ti.drawTH()
	ti.drawTR()

	ti.saveFile()
}
