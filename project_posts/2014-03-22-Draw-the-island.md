I looking to produce and images base on the shape of the Montreal Island.
I want to manualy draw the shape of the Island and then digitilize the draw
and than fill the inside of the island with color, curve base on matrix retrieve from
the number of tweet of the Montreal region.

For now I have tweeter im my mind but a can change my mind after I'll explore different api i'll find
on the web.

Below is the picture of my draw


![Example Image](../project_images/island.jpg?raw=true "Montreal island hand draw")

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
