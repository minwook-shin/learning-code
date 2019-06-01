from PIL import Image
from PIL import ImageFilter


try:
    with Image.open("original.png") as im:
        out = im.filter(ImageFilter.BoxBlur(20))
        out.save("filter_img.png", "PNG")
except IOError:
    pass