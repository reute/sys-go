package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

type Params struct {
	width, height, maxIter int
}

func main() {
	width := flag.Int("w", 800, "width of the image")
	height := flag.Int("h", 800, "height of the image")
	maxIter := flag.Int("i", 1000, "maximum number of iterations")
	flag.Parse()

	params := Params{*width, *height, *maxIter}
	img := GenerateMandelbrot(params)
	SaveImage(img, "mandelbrot.png")
}

func GenerateMandelbrot(p Params) *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, p.width, p.height))

	xMin, xMax := -2.5, 1.5
	yMin, yMax := -1.5, 1.5

	for x := 0; x < p.width; x++ {
		for y := 0; y < p.height; y++ {
			cx := xMin + (xMax-xMin)*float64(x)/float64(p.width)
			cy := yMin + (yMax-yMin)*float64(y)/float64(p.height)
			c := complex(cx, cy)
			color := Mandelbrot(c, p.maxIter)
			img.Set(x, y, color)
		}
	}

	return img
}

func Mandelbrot(c complex128, maxIter int) color.RGBA {
	z := complex(0, 0)
	var n int
	for ; n < maxIter; n++ {
		if cmplx.Abs(z) > 2 {
			break
		}
		z = z*z + c
	}

	if n == maxIter {
		return color.RGBA{0, 0, 0, 255} // Points in the Mandelbrot set are colored black
	}

	// Color gradient based on the number of iterations
	hue := uint8(255 * n / maxIter)
	saturation := uint8(255)
	value := uint8(255 - 255*float64(n)/float64(maxIter))

	r, g, b := hsvToRgb(hue, saturation, value)
	return color.RGBA{r, g, b, 255}
}

func hsvToRgb(h, s, v uint8) (r, g, b uint8) {
	hh := float64(h) / 256.0 * 6.0
	sv := float64(s) / 255.0
	vv := float64(v) / 255.0

	ii := int(hh)
	ff := hh - float64(ii)
	p := vv * (1.0 - sv)
	q := vv * (1.0 - sv*ff)
	t := vv * (1.0 - sv*(1.0-ff))

	switch ii {
	case 0:
		r, g, b = uint8(vv*255), uint8(t*255), uint8(p*255)
	case 1:
		r, g, b = uint8(q*255), uint8(vv*255), uint8(p*255)
	case 2:
		r, g, b = uint8(p*255), uint8(vv*255), uint8(t*255)
	case 3:
		r, g, b = uint8(p*255), uint8(q*255), uint8(vv*255)
	case 4:
		r, g, b = uint8(t*255), uint8(p*255), uint8(vv*255)
	default:
		r, g, b = uint8(vv*255), uint8(p*255), uint8(q*255)
	}

	return
}

func SaveImage(img *image.RGBA, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}
