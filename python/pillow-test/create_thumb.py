from PIL import Image

try:
    im = Image.open("original.png")
    im.thumbnail((128, 128))
    im.save("original.thumbnail", "JPEG")
except IOError:
    pass
