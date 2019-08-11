package tableimage

import (
	"image"
	"reflect"
	"testing"
)

func Test_createContext(t *testing.T) {
	type args struct {
		width           int
		height          int
		backgroundColor string
	}
	tests := []struct {
		name string
		args args
		want tableImage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createContext(tt.args.width, tt.args.height, tt.args.backgroundColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tableImage_setRgba(t *testing.T) {
	type fields struct {
		width           int
		height          int
		backgroundColor string
		img             *image.RGBA
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ti := &tableImage{
				width:           tt.fields.width,
				height:          tt.fields.height,
				backgroundColor: tt.fields.backgroundColor,
				img:             tt.fields.img,
			}
			ti.setRgba()
		})
	}
}

func Test_tableImage_addString(t *testing.T) {
	type fields struct {
		width           int
		height          int
		backgroundColor string
		img             *image.RGBA
	}
	type args struct {
		x     int
		y     int
		label string
		color string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ti := &tableImage{
				width:           tt.fields.width,
				height:          tt.fields.height,
				backgroundColor: tt.fields.backgroundColor,
				img:             tt.fields.img,
			}
			ti.addString(tt.args.x, tt.args.y, tt.args.label, tt.args.color)
		})
	}
}

func Test_tableImage_addLine(t *testing.T) {
	type fields struct {
		width           int
		height          int
		backgroundColor string
		img             *image.RGBA
	}
	type args struct {
		x1    int
		y1    int
		x2    int
		y2    int
		color string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ti := &tableImage{
				width:           tt.fields.width,
				height:          tt.fields.height,
				backgroundColor: tt.fields.backgroundColor,
				img:             tt.fields.img,
			}
			ti.addLine(tt.args.x1, tt.args.y1, tt.args.x2, tt.args.y2, tt.args.color)
		})
	}
}

func Test_tableImage_save(t *testing.T) {
	type fields struct {
		width           int
		height          int
		backgroundColor string
		img             *image.RGBA
	}
	type args struct {
		fileType string
		filePath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ti := &tableImage{
				width:           tt.fields.width,
				height:          tt.fields.height,
				backgroundColor: tt.fields.backgroundColor,
				img:             tt.fields.img,
			}
			ti.save(tt.args.fileType, tt.args.filePath)
		})
	}
}
