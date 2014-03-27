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
	From       int
	To         int
}

type tomlConfig struct {
	Countries map[string]Country
}

func main() {
	var angloSaxon tomlConfig
	if _, err := toml.DecodeFile("settings.toml", &angloSaxon); err != nil {
		fmt.Println(err)
		return
	}

	for name, can := range angloSaxon.Countries {
		fmt.Printf("Server: %s %f \n", name, can.Deathbygun)
	}

	GunsOverview(angloSaxon.Countries, 1000)

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

func GunsOverview(angloSaxon map[string]Country, population int) {
	im, gc := initGc(2250, 1600)
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

	for x := 0.0; x < 2250.0; {

		gc.MoveTo(x, 0.0)
		r, g, b, a := random(0, 245), random(0, 255), random(0, 255), 255

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
	for _, pays := range angloSaxon {
		for x := pays.From; x < pays.To; {
			var curtainsize = pays.Guns * 16
			var curtainsizepos = float64(pays.Rank)
			gc.SetStrokeColor(color.RGBA{255, 255, 255, 240})
			gc.MoveTo(float64(x), curtainsizepos)
			x++
			gc.LineTo(float64(x), curtainsize)
			gc.SetLineWidth(1)
			gc.Stroke()

		}
	}

	gc.SetFontData(draw2d.FontData{"pshift", draw2d.FontFamilyMono, draw2d.FontStyleBold | draw2d.FontStyleItalic})
	for name, pays := range angloSaxon {
		deathbygun, guns, rank := prepareData(pays, population)
		/*
			for j := 1; j < guns; j++ {
				gc.SetLineWidth(0)
				x4, y4 := float64(random(pays.From, pays.To)), float64(random(0, 1600))
				draw2d.Circle(gc, x4, y4, 2)
				gc.SetFillColor(color.RGBA{255, 0, 0, uint8(255)})
				gc.FillStroke()
			}
		*/
		for j := 1; j < deathbygun; j++ {
			gc.SetLineWidth(0)
			x4, y4 := float64(random(pays.From, pays.To)), float64(random(0, 1600))
			draw2d.Circle(gc, x4, y4, 6)
			gc.SetFillColor(color.RGBA{204, 0, 0, 255})
			gc.FillStroke()
		}
		fmt.Printf("name=%d d=%d g=%d r=%d p=%d \n", name, deathbygun, guns, rank)

		gc.SetFillColor(color.Black)
		gc.SetFontSize(120)
		//gc.FillString(name)
		gc.FillStringAt(name, float64(pays.To-360), 1600/2)
		gc.FillString(name)
	}
	for name, pays := range angloSaxon {
		//gc.SetFillColor(color.RGBA{0, 0, 0, 255})
		if name != "UK" {
			gc.SetStrokeColor(color.White)
			gc.MoveTo(float64(pays.To), 0.0)
			gc.LineTo(float64(pays.To), 1600)
			gc.SetLineWidth(25)
			gc.Stroke()
		}
		/*
			gc.SetStrokeColor(color.White)
			gc.MoveTo(float64(pays.From), 0.0)
			gc.LineTo(float64(pays.From), pays.Guns*16)
			gc.SetLineWidth(75)
			gc.Stroke()
			fmt.Printf("guns %d", pays.Guns*16, name)
		*/
	}
	/*
		deathbygun, guns, rank := prepareData(pays["canada"], population)

		for j := 1; j < guns; j++ {
			gc.SetLineWidth(0)
			x4, y4 := float64(random(0, 844)), float64(random(0, 1600))
			draw2d.Circle(gc, x4, y4, 2) // left eye
			gc.SetFillColor(color.RGBA{255, 0, 0, uint8(255)})
			gc.FillStroke()
		}

		for j := 1; j < deathbygun; j++ {
			gc.SetLineWidth(0)
			x4, y4 := float64(random(0, 844)), float64(random(0, 1600))
			draw2d.Circle(gc, x4, y4, 6) // left eye
			gc.SetFillColor(color.Black)
			gc.FillStroke()
		}
		fmt.Printf("canada d=%d g=%d r=%d p=%d \n", deathbygun, guns, rank)

		deathbygun, guns, rank = prepareData(pays["us"], population)

		for j := 1; j < guns; j++ {
			gc.SetLineWidth(0)
			x4, y4 := float64(random(859, 1703)), float64(random(0, 1600))
			draw2d.Circle(gc, x4, y4, 2) // left eye
			gc.SetFillColor(color.RGBA{255, 0, 0, uint8(255)})
			gc.FillStroke()
		}

		for j := 1; j < deathbygun; j++ {
			gc.SetLineWidth(0)
			x4, y4 := float64(random(859, 1703)), float64(random(0, 1600))
			draw2d.Circle(gc, x4, y4, 6) // left eye
			gc.SetFillColor(color.Black)
			gc.FillStroke()
		}
		fmt.Printf("us d=%d g=%d r=%d \n", deathbygun, guns, rank)
		deathbygun, guns, rank = prepareData(pays["uk"], population)

		for j := 1; j < guns; j++ {
			gc.SetLineWidth(0)
			x5, y5 := float64(random(1718, 2562)), float64(random(0, 1600))
			draw2d.Circle(gc, x5, y5, 2) // left eye
			gc.SetFillColor(color.RGBA{255, 0, 0, uint8(255)})
			gc.FillStroke()
		}

		for j := 1; j < deathbygun; j++ {
			gc.SetLineWidth(0)
			x6, y6 := float64(random(1718, 2562)), float64(random(0, 1600))
			draw2d.Circle(gc, x6, y6, 6) // left eye
			gc.SetFillColor(color.Black)
			gc.FillStroke()
		}
		fmt.Printf("UK d=%d g=%d r=%d \n", deathbygun, guns, rank)
		x, y := 1.0, 300.0
		x1, y1 := 300.4, 400.4
		x2, y2 := 700.6, 200.6
		x3, y3 := 945.4, 300.0
		gc.SetFillColor(color.RGBA{156, 9, 29, 255})
		gc.SetStrokeColor(color.RGBA{156, 6, 29, 255})
		gc.SetLineWidth(10)
		gc.MoveTo(x, y)
		gc.CubicCurveTo(x1, y1, x2, y2, x3, y3)
		gc.Stroke()
		gc.MoveTo(800, 420)
		gc.SetFillColor(color.White)
		gc.SetFontSize(120)
		gc.SetFontData(draw2d.FontData{"eltobito", draw2d.FontFamilyMono, draw2d.FontStyleBold | draw2d.FontStyleItalic})

		gc.FillString("CA")
		gc.FillStringAt("CA", 400.0, 800)
		gc.FillString("CA")

		gc.MoveTo(844, 0.0)
		gc.LineTo(844, 1600)
		gc.SetLineWidth(30)
		gc.Stroke()

		gc.FillString("US")
		gc.FillStringAt("US", 1200.0, 800)
		gc.FillString("US")

		gc.MoveTo(1715, 0.0)
		gc.LineTo(1715, 1600)
		gc.SetLineWidth(30)
		gc.Stroke()

		gc.FillString("UK")
		gc.FillStringAt("UK", 2044.0, 800)
		gc.FillString("UK")
		gc.Stroke()*/

	saveToPngFile("GunsOverview", im)
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
