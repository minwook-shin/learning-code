from PIL import Image

try:
    with Image.open("original.png") as im:
        print(im.format, im.size, im.mode)
except IOError:
    pass
