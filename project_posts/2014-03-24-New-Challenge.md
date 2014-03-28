I decide to use another language instead of Python. The reason is that  I won't need to
use a Google api to be respect the rules of the project. I have done a couple of research and 
I saw that GO have an image library. So I decide to learn and use that language from Google
and I won't have any string attached for the api. Oups I forgot...I don't know GO...So I have 1 week to know how to use it
and be able to deliver the project.


I have done a couple of test to generate an image where every pixels are random. 
Specifically I tried to reproduce the test I have done in python but with GO language.
 



![Example Image](../project_images/go.png?raw=true "Random pixel")

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

          x++
          gc.LineTo(x, y)
          //gc.LineTo(x3, y3)
          gc.Stroke()
        }
      }
    
      saveToPngFile("TestDrawMosaic", im)
    }

