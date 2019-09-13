package tableimage

import (
	"image/jpeg"
	"image/png"
	"os"
)

func (ti *tableImage) drawTH() {
	for colNo, td := range ti.th.Tds {
		ti.addString(colNo*columnSpace+tablePadding, 1*rowSpace, td.Text, td.Color)
		ti.addLine(colNo*columnSpace, 0, colNo*columnSpace, ti.height, "#000000")
	}
	//draw the end horizontal line
	ti.addLine(ti.width-1, 0, ti.width-1, ti.height, "#000000")
	//draw the start horizontal line
	ti.addLine(1, 1, ti.width, 1, "#000000")
	//draw the double line to signal it is a th
	ti.addLine(1, 1*rowSpace+separatorPadding, ti.width, 1*rowSpace+separatorPadding, "#000000")
	ti.addLine(1, 1*rowSpace+separatorPadding+2, ti.width, 1*rowSpace+separatorPadding+2, "#000000")

}

func (ti *tableImage) drawTR() {
	fRowNo := 2
	for _, tds := range ti.trs {
		//start with the second row since the first one is the th
		maxRowHeight := 1
		for colNo, td := range tds.Tds {
			wrapedTexts := wrapText(td.Text)
			noRowPerCol := len(wrapedTexts)
			if maxRowHeight < noRowPerCol {
				maxRowHeight = noRowPerCol
			}
			for noTextRows, wrapedText := range wrapedTexts {
				colHeight := (fRowNo + noTextRows) * rowSpace
				colWidth := colNo*columnSpace + tablePadding

				if colNo == 0 {
					colWidth = columnSpace - columnSpace + tablePadding
				}

				ti.addString(colWidth, colHeight, wrapedText, td.Color)
			}
		}
		fRowNo += maxRowHeight
		ti.addLine(1, (fRowNo-1)*rowSpace+separatorPadding, ti.width, (fRowNo-1)*rowSpace+separatorPadding, "#000")
	}
}

func (ti *tableImage) calculateHeight() {
	//start from 1 since we have th
	totalRowNo := 1
	for _, tr := range ti.trs {
		maxRowHeight := 1
		for _, td := range tr.Tds {
			wrapedText := wrapText(td.Text)
			// in case we have a multi line text
			if len(wrapedText) > maxRowHeight {
				maxRowHeight = len(wrapedText)
			}
		}
		totalRowNo += maxRowHeight
	}
	ti.height = totalRowNo*rowSpace + rowSpace - tablePadding + 5
}

func (ti *tableImage) calculateWidth() {
	totalColumnNo := len(ti.th.Tds)

	ti.width = totalColumnNo * columnSpace
}

func (ti *tableImage) saveFile() {
	f, err := os.Create(ti.filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	if ti.fileType == "jpg" {
		if err := jpeg.Encode(f, ti.img, nil); err != nil {
			panic(err)
		}
	} else {
		if err := png.Encode(f, ti.img); err != nil {
			panic(err)
		}
	}
}
