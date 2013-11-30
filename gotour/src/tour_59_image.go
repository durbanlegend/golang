package main
 
import (
	"image"
    "code.google.com/p/go-tour/pic"
	"image/color"
)

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	for i := range arr {
		arr[i] = make([]uint8, dx)
		for j := range arr[i] {
			// arr[i][j] = uint8(i ^ j)
			// arr[i][j] = uint8((i + j)/2)
            arr[i][j] = uint8(i * j)
		}
	}
	return arr
}

type Image struct{
	Width, Height int
	colr uint8	
}
 
func (r *Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, r.Width, r.Height)
}
 
func (r *Image) ColorModel() color.Model {
	return color.RGBAModel
}
 
func (r *Image) At(x, y int) color.Color {
	return color.RGBA{r.colr+uint8(x ^ y), r.colr+uint8(x * y), r.colr+uint8((x + y) / 5), 255}
}
 
func main() {
	m := Image{1020, 255, 0}
	pic.ShowImage(&m)
}
 