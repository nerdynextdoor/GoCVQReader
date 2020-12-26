package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
	"math"
	"time"

	"github.com/llgcode/draw2d/draw2dimg"
	"gonum.org/v1/plot"

	"./QRreader/goqr"
)

func recognizeFile(path string) {
	start := time.Now()
	fmt.Printf("recognize file: %v\n", path)
	imgdata, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	img, _, err := image.Decode(bytes.NewReader(imgdata))
	if err != nil {
		fmt.Printf("image.Decode error: %v\n", err)
		return
	}
	qrCodes, err := goqr.Recognize(img)
	if err != nil {
		fmt.Printf("Recognize failed: %v\n", err)
		return
	}

	for _, qrCode := range qrCodes {

		fmt.Printf("qrCode text: %s\n", qrCode.Payload)

	}
	elapsed := time.Since(start)
	log.Printf("Time taken to find QRcode: %s", elapsed)
	fmt.Printf("QR corner location:  %v \n", goqr.GV)

	//p1.x
	angleTest(goqr.Gx1, goqr.Gy1, goqr.Gx2, goqr.Gy2)
	//angleTest(goqr.GV[0],)
	//fmt.Printf("image.Decode error: %v\n", goqr.)
}

func main() {
	recognizeFile("QRgrid.png")

	//drawing2()
}

func angleTest(x1, y1, x2, y2 float64) {
	ydiff := y2 - y1
	xdiff := x2 - x1
	radians := math.Atan2(ydiff, xdiff)
	//fmt.Printf("Ydiff: %v\n", ydiff)
	//fmt.Printf("Xdiff: %v\n", xdiff)
	//fmt.Printf("radians: %v\n", radians)

	angle := int(radians * 180 / math.Pi)

	fmt.Printf("QR Rotation: %v degrees\n", angle)

}

func drawing2() {
	// Initialize the graphic context on an RGBA image
	dest := image.NewRGBA(image.Rect(0, 0, 600, 600))
	gc := draw2dimg.NewGraphicContext(dest)

	// Set some properties
	gc.SetFillColor(color.RGBA{0x44, 0xff, 0x44, 0xff})
	gc.SetStrokeColor(color.RGBA{0x44, 0x44, 0x44, 0xff})
	gc.SetLineWidth(5)

	// Draw a closed shape
	gc.MoveTo(314, 258) // should always be called first for a new path
	//gc.LineTo(342, 258)
	//gc.QuadCurveTo(341, 258, 314, 285)
	gc.CubicCurveTo(341, 258, 341, 385, 314, 258)
	gc.Close()
	gc.FillStroke()

	// Save to file
	draw2dimg.SaveToPngFile("./hello.png", dest)
}

func drawing() {
	p, err := plot.New()
	if err != nil {
		log.Panic(err)
	}

	//plotter.DefaultLineStyle.Width = vg.Points(1)
	//plotter.DefaultGlyphStyle.Radius = vg.Points(3)

	//p.Y.Tick.Marker = plot.ConstantTicks([]plot.Tick{
	//	{Value: 0, Label: "0"}, {Value: 0.25, Label: ""}, {Value: 0.5, Label: "0.5"}, {Value: 0.75, Label: ""}, {Value: 1, Label: "1"},
	//})
	//p.X.Tick.Marker = plot.ConstantTicks([]plot.Tick{
	//	{Value: 0, Label: "0"}, {Value: 0.25, Label: ""}, {Value: 0.5, Label: "0.5"}, {Value: 0.75, Label: ""}, {Value: 1, Label: "1"},
	//})

	//pts := plotter.XYs{{X: 0, Y: 0}, {X: 0, Y: 1}, {X: 0.5, Y: 1}, {X: 0.5, Y: 0.6}, {X: 0, Y: 0.6}}
	//line, err := plotter.NewLine(pts)
	//if err != nil {
	//	log.Panic(err)
	//}
	//scatter, err := plotter.NewScatter(pts)
	//if err != nil {
	//	log.Panic(err)
	//}
	//p.Add(line, scatter)

	//pts = plotter.XYs{{X: 1, Y: 0}, {X: 0.75, Y: 0}, {X: 0.75, Y: 0.75}}
	//line, err = plotter.NewLine(pts)
	//if err != nil {
	//	log.Panic(err)
	//}
	//scatter, err = plotter.NewScatter(pts)
	//if err != nil {
	//	log.Panic(err)
	//}
	//p.Add(line, scatter)

	//pts = plotter.XYs{{X: 0.5, Y: 0.5}, {X: 1, Y: 0.5}}
	//line, err = plotter.NewLine(pts)
	//if err != nil {
	//	log.Panic(err)
	//}
	//scatter, err = plotter.NewScatter(pts)
	//if err != nil {
	//	log.Panic(err)
	//}
	//p.Add(line, scatter)

	err = p.Save(100, 100, "./plotLogo.png")
	if err != nil {
		log.Panic(err)
	}
}
