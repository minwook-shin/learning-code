from PIL import Image


try:
    with Image.open("original.png") as im:
        box = (100, 100, 400, 400)
        region = im.crop(box)
        region = region.transpose(Image.ROTATE_180)
        im.paste(region, box)
        im.save("pasted_img.png", "PNG")
        print(region)
except IOError:
    pass