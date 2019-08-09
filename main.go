package main

import (
	"github.com/fogleman/gg"
)

var ths = []string{
	"Name",
	"Duration",
	"Errors No",
	"Fetched Rows",
	"Generated Urls",
	"Uploaded to bucket",
	"Started At",
	"Ended At",
}

var trs = [][]string{
	{
		"hotelreviews",
		"3m4s",
		"200",
		"150000",
		"450000",
		"1800000",
		"2019-08-09 10:16:29",
		"2019-08-09 10:16:29",
	},
	{
		"hotelreviews",
		"3m4s",
		"200",
		"150000",
		"450000",
		"1800000",
		"2019-08-09 10:16:29",
		"2019-08-09 10:16:29",
	},
	{
		"hotelreviews",
		"3m4s",
		"200",
		"150000",
		"450000",
		"1800000",
		"2019-08-09 10:16:29",
		"2019-08-09 10:16:29",
	},
	{
		"hotelreviews",
		"3m4s",
		"200",
		"150000",
		"450000",
		"1800000",
		"2019-08-09 10:16:29",
		"2019-08-09 10:16:29",
	},
}

const (
	rowSpace         = 30
	tablePadding     = 20
	letterPerPx      = 10
	separatorPadding = 10
)

func main() {

	var totalRowNo = len(trs) + 1
	var totalColumnNo = len(ths)
	var columnSpace = getMaxAmountOfLetters() * letterPerPx
	var H = totalRowNo*rowSpace + rowSpace
	var W = totalColumnNo * int(columnSpace)

	dc := gg.NewContext(W, H)

	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)

	for colNo, th := range ths {
		dc.DrawString(th, float64(colNo)*columnSpace+tablePadding, 1*rowSpace)
	}

	//draw the double line for th's
	dc.DrawLine(1, 1*rowSpace+separatorPadding, float64(W), 1*rowSpace+separatorPadding)
	dc.DrawLine(1, 1*rowSpace+separatorPadding+2, float64(W), 1*rowSpace+separatorPadding+2)

	dc.SetRGB255(0, 153, 0)

	for rowNo, tds := range trs {
		//start with the second row since the firsone is the th
		fRowNo := float64(rowNo + 2)
		dc.SetRGB255(255, 0, 0)
		for colNo, td := range tds {
			fColNo := float64(colNo)
			if colNo == 0 {
				dc.DrawString(td, columnSpace-columnSpace+tablePadding, fRowNo*rowSpace)
			}
			dc.DrawString(td, fColNo*columnSpace+tablePadding, fRowNo*rowSpace)
		}
		dc.DrawLine(1, fRowNo*rowSpace+separatorPadding, float64(W), fRowNo*rowSpace+separatorPadding)
	}

	//set table line color
	dc.SetRGB(0, 0, 0)
	dc.Stroke()
	dc.SavePNG("out.png")
}

func getMaxAmountOfLetters() float64 {
	var maxColSpace = 0
	for _, th := range ths {
		if len(th) > maxColSpace {
			maxColSpace = len(th)
		}
	}
	return float64(maxColSpace)
}
