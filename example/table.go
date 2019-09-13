package main

import (
	"tableimage"
)

func main() {
	ti := tableimage.Init("#fff", tableimage.PNG, "./test.png")

	ti.AddTH(
		tableimage.TR{
			BorderColor: "#000",
			Tds: []tableimage.TD{
				tableimage.TD{
					Color: "#000",
					Text:  "Id",
				},
				tableimage.TD{
					Color: "#000",
					Text:  "Name",
				},
				tableimage.TD{
					Color: "#008000",
					Text:  "Price",
				},
			},
		},
	)

	ti.AddTRs(
		[]tableimage.TR{
			tableimage.TR{
				BorderColor: "#000",
				Tds: []tableimage.TD{
					tableimage.TD{
						Color: "#000",
						Text:  "2223",
					},
					tableimage.TD{
						Color: "#000",
						Text:  "Really cool product on two lines",
					},
					tableimage.TD{
						Color: "#0000ff",
						Text:  "2000$",
					},
				},
			},
			tableimage.TR{
				BorderColor: "#000",
				Tds: []tableimage.TD{
					tableimage.TD{
						Color: "#000",
						Text:  "11",
					},
					tableimage.TD{
						Color: "#000",
						Text:  "A more cooler product this time on 3 lines",
					},
					tableimage.TD{
						Color: "#0000ff",
						Text:  "200$",
					},
				},
			},
			tableimage.TR{
				BorderColor: "#000",
				Tds: []tableimage.TD{
					tableimage.TD{
						Color: "#000",
						Text:  "2231",
					},
					tableimage.TD{
						Color: "#000",
						Text:  "Lenovo",
					},
					tableimage.TD{
						Color: "#000",
						Text:  "20400$",
					},
				},
			},
		},
	)
	ti.Save()
}
