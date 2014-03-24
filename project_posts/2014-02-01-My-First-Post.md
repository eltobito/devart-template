After exploring possibility with image processing I install PIL package. And generate a random image to see the result and start working with this snippet of code

    import Image,ImageDraw
    from random import randint as rint
    
    def randgradient():
        img = Image.new("RGB", (300,300), "#FFFFFF")
        draw = ImageDraw.Draw(img)
    
        r,g,b = rint(0,255), rint(0,255), rint(0,255)
        dr = (rint(0,255) - r)/300.
        dg = (rint(0,255) - g)/300.
        db = (rint(0,255) - b)/300.
        for i in range(300):
            r,g,b = r+dr, g+dg, b+db
            draw.line((i,0,i,300), fill=(int(r),int(g),int(b)))
    
        img.save("out.png", "PNG")
    
    if __name__ == "__main__":
        randgradient()

![Example Image](../project_images/out.png?raw=true "Example Image")


