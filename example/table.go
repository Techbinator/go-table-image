package main

import (
	tableimage "github.com/Techbinator/go-table-image"
)

func main() {
	ti := tableimage.Init("#fff", tableimage.PNG, "./test.png")

	ti.AddTH(
		tableimage.TR{
			BorderColor: "#000",
			Tds: []tableimage.TD{
				{
					Color: "#000",
					Text:  "Id",
				},
				{
					Color: "#000",
					Text:  "Name",
				},
				{
					Color: "#008000",
					Text:  "Price",
				},
			},
		},
	)

	ti.AddTRs(
		[]tableimage.TR{
			{
				BorderColor: "#000",
				Tds: []tableimage.TD{
					{
						Color: "#000",
						Text:  "2223",
					},
					{
						Color: "#000",
						Text:  "Really cool product on two lines",
					},
					{
						Color: "#0000ff",
						Text:  "2000$",
					},
				},
			},
			{
				BorderColor: "#000",
				Tds: []tableimage.TD{
					{
						Color: "#000",
						Text:  "11",
					},
					{
						Color: "#000",
						Text:  "A more cooler product this time on 3 lines",
					},
					{
						Color: "#0000ff",
						Text:  "200$",
					},
				},
			},
			{
				BorderColor: "#000",
				Tds: []tableimage.TD{
					{
						Color: "#000",
						Text:  "2231",
					},
					{
						Color: "#000",
						Text:  "Lenovo",
					},
					{
						Color: "#000",
						Text:  "20400$",
					},
				},
			},
		},
	)
	ti.Save()
}
