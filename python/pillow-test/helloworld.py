from PIL import Image

im = Image.open("original.png")
print(im.format, im.size, im.mode)
im.show()