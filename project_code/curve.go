package main

import (
	"bufio"
	"code.google.com/p/draw2d/draw2d"
	"fmt"
	"github.com/BurntSushi/toml"
	"image"
	"image/color"

	"image/png"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	w, h = 512, 512
)

var (
	lastTime int64
	folder   = "../resource/result/"
)

type Country struct {
	Deathbygun float64
	Guns       float64
	Rank       int
	Pop        int
}

type tomlConfig struct {
	Countries map[string]Country
}

func main() {
	var pays tomlConfig
	if _, err := toml.DecodeFile("settings.toml", &pays); err != nil {
		fmt.Println(err)
		return
	}

	for name, can := range pays.Countries {
		fmt.Printf("Server: %s %f \n", name, can.Deathbygun)
	}

	//TestStar()
	//var pop = averagepop(pays.Countries["canada"].Pop, pays.Countries["us"].Pop, pays.Countries["uk"].Pop)

	TestDrawLines(pays.Countries, 1000)
	//TestDrawMosaic()
	//TestFillString()
	//spot()

}

func Show(name string) {
	command := "open"
	arg1 := "-a"
	arg2 := "/Applications/Preview.app"
	cmd := exec.Command(command, arg1, arg2, name)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func TestDrawLines(pays map[string]Country, population int) {
	im, gc := initGc(2586, 1600)
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("population %d \n", population)
	// draw a cubic curve

	//r, g, b, a := random(0, 255), random(0, 255), random(0, 255), 255

	//fmt.Printf("%d\nallo%d", b,r)
	//b1 := m.Bounds()
	gc.SetLineWidth(1)
	gc.MoveTo(1.0, 1.0)

	//gc.SetStrokeColor(cstart)
	//gc.Stroke()
	//for y := 1.0; y < 600.0; y++ {

	for x := 0.0; x < 2559.0; {

		gc.MoveTo(x, 0.0)
		r, g, b, a := rand.Intn(200), rand.Intn(200), rand.Intn(200), 255

		//rg := random(0, rank)
		//r, g, b, a := r, g, b, rank

		cstart := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		gc.SetStrokeColor(cstart)

		// draw segment of curve

		x++
		gc.LineTo(x, 1600)
		//gc.LineTo(x3, y3)
		gc.Stroke()
	}
	deathbygun, guns, rank := prepareData(pays["canada"], population)

	for j := 1; j < guns; j++ {
		gc.SetLineWidth(0)
		x4, y4 := float64(random(10, 844)), float64(random(0, 1600))
		draw2d.Circle(gc, x4, y4, 2) // left eye
		gc.SetFillColor(color.RGBA{255, 255, 0, uint8(255)})
		gc.FillStroke()
	}

	for j := 1; j < deathbygun; j++ {
		gc.SetLineWidth(0)
		x4, y4 := float64(random(10, 844)), float64(random(0, 1600))
		draw2d.Circle(gc, x4, y4, 6) // left eye
		gc.SetFillColor(color.Black)
		gc.FillStroke()
	}
	fmt.Printf("canada d=%d g=%d r=%d p=%d \n", deathbygun, guns, rank)

	deathbygun, guns, rank = prepareData(pays["us"], population)

	for j := 1; j < guns; j++ {
		gc.SetLineWidth(0)
		x4, y4 := float64(random(871, 1715)), float64(random(0, 1600))
		draw2d.Circle(gc, x4, y4, 2) // left eye
		gc.SetFillColor(color.White)
		gc.FillStroke()
	}

	for j := 1; j < deathbygun; j++ {
		gc.SetLineWidth(0)
		x4, y4 := float64(random(871, 1715)), float64(random(0, 1600))
		draw2d.Circle(gc, x4, y4, 6) // left eye
		gc.SetFillColor(color.Black)
		gc.FillStroke()
	}
	fmt.Printf("us d=%d g=%d r=%d \n", deathbygun, guns, rank)
	deathbygun, guns, rank = prepareData(pays["uk"], population)

	for j := 1; j < guns; j++ {
		gc.SetLineWidth(0)
		x5, y5 := float64(random(1742, 2586)), float64(random(0, 1600))
		draw2d.Circle(gc, x5, y5, 2) // left eye
		gc.SetFillColor(color.RGBA{255, 0, 0, uint8(255)})
		gc.FillStroke()
	}

	for j := 1; j < deathbygun; j++ {
		gc.SetLineWidth(0)
		x6, y6 := float64(random(1742, 2586)), float64(random(0, 1600))
		draw2d.Circle(gc, x6, y6, 6) // left eye
		gc.SetFillColor(color.Black)
		gc.FillStroke()
	}
	fmt.Printf("UK d=%d g=%d r=%d \n", deathbygun, guns, rank)
	x, y := 1.0, 300.0
	x1, y1 := 300.4, 400.4
	x2, y2 := 700.6, 200.6
	x3, y3 := 945.4, 300.0
	gc.SetFillColor(color.RGBA{255, 128, 0, 255})
	gc.SetStrokeColor(color.RGBA{255, 128, 0, 255})
	gc.SetLineWidth(10)
	gc.MoveTo(x, y)
	gc.CubicCurveTo(x1, y1, x2, y2, x3, y3)
	gc.Stroke()
	gc.MoveTo(844, 0.0)
	gc.LineTo(844, 1600)
	gc.SetLineWidth(27)
	gc.Stroke()
	gc.MoveTo(1715, 0.0)
	gc.LineTo(1715, 1600)
	gc.SetLineWidth(27)
	gc.Stroke()

	saveToPngFile("TestDrawLines", im)
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func prepareData(pays Country, pop int) (int, int, int) {

	var deathbygun = pays.Deathbygun * float64(pop)
	var guns = pays.Guns * float64(pop)
	var rank = int(pays.Rank * 255 / 178)

	return int(deathbygun), int(guns), rank
}

func canada() (int, int, int, int) {
	var deathbygun = 0.5 * 50000 / 100
	var guns = 23.8 * 50000 / 100
	var rank = int(12 * 255 / 178)
	var pop = 35344962
	return int(deathbygun), int(guns), rank, pop
}

func us() (int, int, int, int) {
	var deathbygun = 3.6 * 50000 / 100
	var guns = 100.0 * 50000 / 100
	var rank = int(1 * 255 / 178)
	var pop = 317751000
	return int(deathbygun), int(guns), rank, pop

}

func uk() (int, int, int, int) {
	var deathbygun = 0.06 * 50000 / 100
	var guns = 6.7 * 50000 / 100
	var rank = int(22 * 255 / 178)
	var pop = 63705000
	return int(deathbygun), int(guns), rank, pop
}

func averagepop(popca, popus, popuk int) int {
	avr_pop := (popca + popus + popuk) / 3
	return int(avr_pop)
}

func initGc(w, h int) (image.Image, draw2d.GraphicContext) {
	i := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2d.NewGraphicContext(i)

	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(image.White)

	// fill the background
	//gc.Clear()

	return i, gc
}

func saveToPngFile(TestName string, m image.Image) {
	filePath := TestName + ".png"
	f, err := os.Create(filePath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer f.Close()
	b := bufio.NewWriter(f)
	err = png.Encode(b, m)
	if err != nil {

		log.Println(err)
		os.Exit(1)
	}
	err = b.Flush()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Wrote %s OK.\n", filePath)
	Show(filePath)
}
func TestDrawCubicCurve() {
	i, gc := initGc(w, h)
	// draw a cubic curve
	x, y := 25.6, 128.0
	x1, y1 := 102.4, 230.4
	x2, y2 := 153.6, 25.6
	x3, y3 := 230.4, 128.0
	gc.SetFillColor(color.RGBA{0, 0, 255, 255})
	gc.SetLineWidth(10)
	gc.MoveTo(x, y)
	gc.CubicCurveTo(x1, y1, x2, y2, x3, y3)
	gc.Stroke()

	gc.SetStrokeColor(color.RGBA{0, 200, 255, 255})

	gc.SetLineWidth(6)
	// draw segment of curve
	gc.MoveTo(x, y)
	gc.LineTo(x1, y1)
	gc.LineTo(x2, y2)
	gc.LineTo(x3, y3)
	gc.Stroke()
	saveToPngFile("TestDrawCubicCurve", i)
}
func TestStar() {
	i, gc := initGc(w, h)

	// draw a cubic curve
	rand.Seed(time.Now().UnixNano())
	rand.Intn(255)
	r, g, b, a := rand.Intn(255), rand.Intn(255), rand.Intn(255), 255
	var (
		cstart color.Color = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	)

	for i := 0.0; i < 360; i++ { // Go from 0 to 360 degrees in 10 degree steps
		gc.Save()
		r, g, b, a := rand.Intn(255), rand.Intn(255), rand.Intn(255), 255
		//r, g, b, a := 255, 128, 0, rand.Intn(255)
		cstart = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		gc.SetStrokeColor(cstart)
		gc.SetLineWidth(1) // Keep rotations temporary
		gc.Translate(144, 144)
		gc.Rotate(i * (math.Pi / 180.0)) // Rotate by degrees on stack from 'for'
		gc.MoveTo(0, 0)
		gc.LineTo(20, 0)
		gc.Stroke()
		gc.Restore()
	}
	saveToPngFile("TestStar", i)
}
func TestFillString() {
	i, gc := initGc(300, 300)
	r, g, b, a := 255, 0, 0, 255
	var (
		cstart = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	)
	draw2d.RoundRect(gc, 5, 5, 95, 95, 10, 10)
	//draw2d.Circle(path, cx, cy, radius)
	gc.SetFillColor(cstart)
	gc.FillStroke()
	gc.SetFontSize(25)
	gc.MoveTo(10, 52)
	gc.SetFillColor(color.White)
	gc.SetFontData(draw2d.FontData{"eltobito", draw2d.FontFamilyMono, draw2d.FontStyleBold | draw2d.FontStyleItalic})
	width := gc.FillString("cou")
	fmt.Printf("width: %f\n", width)
	gc.RMoveTo(width+1, 0)
	gc.FillString("cou")
	gc.FillStringAt("Tobie Desjardins", 5.0, 150)
	gc.FillString("Tobie Desjardins")
	saveToPngFile("TestFillString", i)
}

/*
func spot2(i image) {
        gc := draw2d.NewGraphicContext(i)
        for j := 1; j < 100; j++ {
                gc.SetLineWidth(1)
                x4, y4 := float64(random(325, 605)), float64(random(10, 585))
                draw2d.Circle(gc, x4, y4, 2) // left eye
                gc.SetFillColor(color.White)
                gc.FillStroke()
        }

        for j := 1; j < 10; j++ {
                gc.SetLineWidth(1)
                x4, y4 := float64(random(325, 605)), float64(random(10, 585))
                draw2d.Circle(gc, x4, y4, 2) // left eye
                gc.SetFillColor(color.Black)
                gc.FillStroke()
        }
*/
func spot() {
	x, y := 50.0, 50.0
	im, gc := initGc(300, 300)

	draw2d.Circle(gc, x+60, y+45, 60) // left eye
	gc.SetFillColor(color.Black)
	gc.FillStroke()

	saveToPngFile("spot", im)
}

func TestDrawMosaic() {
	im, gc := initGc(300, 300)
	// draw a cubic curve
	rand.Seed(time.Now().UnixNano())
	rand.Intn(255)

	r, g, b, a := random(0, 255), random(0, 255), random(0, 255), 255

	//fmt.Printf("%d\nallo%d", b,r)
	//b1 := m.Bounds()
	gc.SetLineWidth(1)
	gc.MoveTo(1.0, 1.0)
	var (
		cstart color.Color = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
	)
	gc.SetStrokeColor(cstart)
	gc.Stroke()
	for y := 1.0; y < 300.0; y++ {
		for x := 1.0; x < 300.0; {

			gc.MoveTo(x, y)
			r, g, b, a := random(0, 255), random(0, 255), random(0, 255), 255
			cstart = color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
			gc.SetStrokeColor(cstart)

			// draw segment of curve

			x++
			gc.LineTo(x, y)
			//gc.LineTo(x3, y3)
			gc.Stroke()
		}
	}

	saveToPngFile("TestDrawMosaic", im)
}
