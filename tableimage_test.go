package tableimage

import (
	"testing"
)

var ths = []string{
	"Name",
	"Duration",
	"Errors No",
	"Fetched Rows",
	"Ended At",
}

var trs = [][]string{
	{
		"hotelreviews",
		"3m4s",
		"200",
		"150000",
		"2019-08-09 10:16:29",
	},
	{
		"hotelreviews",
		"3m4s",
		"200",
		"150000",
		"2019-08-09 10:16:29",
	},
	{
		"hotelreviews",
		"3m4s",
		"200",
		"150000",
		"2019-08-09 10:16:29",
	},
	{
		"hotelreviews",
		"3m4s",
		"200",
		"150000",
		"2019-08-09 10:16:29",
	},
}

func TestGenerateTableImage(t *testing.T) {
	type args struct {
		ths []string
		trs [][]string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "generate table",
			args: args{
				ths: ths,
				trs: trs,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateTableImage(tt.args.ths, tt.args.trs)
		})
	}
}
