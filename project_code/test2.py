import Image
picture = Image.open("picture.jpg")

# Get the size of the image
width, height = picture.size()

# Process every pixel
for x in width:
   for y in height:
       current_color = picture.getpixel( (x,y) )
       print current_color
       ####################################################################
       # Do your logic here and create a new (R,G,B) tuple called new_color
       ####################################################################
       picture.putpixel( (x,y), new_color)
