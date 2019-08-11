package tableimage

func (ti *tableImage) drawTH(ths []string) {
	for colNo, th := range ths {
		ti.addString(colNo*ti.columnSpace+tablePadding, 1*rowSpace, th, "#000000")
	}

	//draw the double line to signal it is a th
	ti.addLine(1, 1*rowSpace+separatorPadding, ti.width, 1*rowSpace+separatorPadding, "#000000")
	ti.addLine(1, 1*rowSpace+separatorPadding+2, ti.width, 1*rowSpace+separatorPadding+2, "#000000")
}

func (ti *tableImage) drawTR(trs [][]string) {
	for rowNo, tds := range trs {
		//start with the second row since the firsone is the th
		fRowNo := rowNo + 2
		for colNo, td := range tds {

			if colNo == 0 {
				ti.addString(ti.columnSpace-ti.columnSpace+tablePadding, fRowNo*rowSpace, td, "#000")
			}
			ti.addString(colNo*ti.columnSpace+tablePadding, fRowNo*rowSpace, td, "#000")
		}
		ti.addLine(1, fRowNo*rowSpace+separatorPadding, ti.width, fRowNo*rowSpace+separatorPadding, "#000")
	}
}
