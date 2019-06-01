from PIL import Image


try:
    with Image.open("original.png") as im:
        out = im.resize((128, 128))
        out.save("resize_img.png", "PNG")
        out = im.rotate(45)  # degrees counter-clockwise
        out.save("rotate_img.png", "PNG")
except IOError:
    pass