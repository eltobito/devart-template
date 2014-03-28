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

type Country struct {
	Deathbygun float64
	Guns       float64
	Rank       int
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

	GunsOverview(angloSaxon.Countries)

}

//work on mac
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

func GunsOverview(angloSaxon map[string]Country) {
	im, gc := initGunControlOverviewArtWork(3200, 2000)
	rand.Seed(time.Now().UnixNano())
	gc.SetLineWidth(1)
	fmt.Printf("Draw background\n")
	for x := 20.0; x < 3180.0; {

		gc.MoveTo(x, 19.0)
		r, g, b, a := random(0, 245), random(0, 255), random(0, 255), 255
		cstart := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
		gc.SetStrokeColor(cstart)

		x++
		gc.LineTo(x, 1981)
		gc.Stroke()
	}
	fmt.Printf("Draw curtain: dependancy gun pocession rank and rate of gun per 100 person\n")
	for _, pays := range angloSaxon {
		for x := pays.From; x < pays.To; {
			var curtainsize = pays.Guns * 1981 / 100
			var curtainsizepos = float64(pays.Rank + 19)
			gc.SetStrokeColor(color.RGBA{255, 255, 255, 240})
			gc.MoveTo(float64(x), curtainsizepos)
			x++
			gc.LineTo(float64(x), curtainsize)
			gc.SetLineWidth(1)
			gc.Stroke()

		}
	}
	fmt.Printf("Gun shot : dependancy gun deah by 100 000 persons\n")
	gc.SetFontData(draw2d.FontData{"pshift", draw2d.FontFamilyMono, draw2d.FontStyleBold | draw2d.FontStyleItalic})
	for name, pays := range angloSaxon {
		deathbygun, guns, rank := prepareData(pays)
		for i := 1; i < deathbygun; i++ {
			gc.SetLineWidth(0)
			x, y := float64(random(pays.From, pays.To)), float64(random(24, 1976))
			draw2d.Circle(gc, x, y, 10)
			gc.SetFillColor(color.RGBA{204, 0, 0, 255})
			gc.FillStroke()
		}
		fmt.Printf("name=%s death=%d guns=%d rank=%d \n", name, deathbygun, guns, rank)
		fmt.Printf("Identify countries\n")
		gc.SetFillColor(color.Black)
		gc.SetFontSize(120)
		gc.FillStringAt(name, float64(pays.To-360), 2000/2)
		gc.FillString(name)
	}
	fmt.Printf("Draw contry separator\n")
	for name, pays := range angloSaxon {
		if name != "UK" {
			gc.SetStrokeColor(color.Black)
			gc.MoveTo(float64(pays.To), 19.0)
			gc.LineTo(float64(pays.To), 1981)
			gc.SetLineWidth(38)
			gc.Stroke()
		}
	}

	saveToPngFile("GunsOverview", im)
}

func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

func prepareData(pays Country) (int, int, int) {
	// 2087974 pixels by country
	var deathbygun = pays.Deathbygun * float64(2087974) / 100000
	var guns = pays.Guns
	var rank = int(pays.Rank * 255 / 178)

	return int(deathbygun), int(guns), rank
}

func initGunControlOverviewArtWork(w, h int) (image.Image, draw2d.GraphicContext) {
	i := image.NewRGBA(image.Rect(0, 0, w, h))
	gc := draw2d.NewGraphicContext(i)
	//Set Frame
	gc.SetStrokeColor(image.Black)
	gc.SetFillColor(image.Black)
	gc.MoveTo(0.0, 0.0)
	gc.LineTo(3200.0, 0.0)
	gc.LineTo(3200.0, 2000.0)
	gc.LineTo(0.0, 2000.0)
	gc.LineTo(0.0, 0.0)
	gc.SetLineWidth(38)

	gc.Stroke()

	return i, gc
}

func saveToPngFile(TestName string, m image.Image) {
	filePath := "../project_images/" + TestName + ".png"
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
