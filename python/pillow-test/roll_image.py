from PIL import Image


try:
    with Image.open("original.png") as im:
        x_size, y_size = im.size

        delta = 300 % x_size

        part1 = im.crop((0, 0, delta, y_size))
        part2 = im.crop((delta, 0, x_size, y_size))
        im.paste(part1, (x_size - delta, 0, x_size, y_size))
        im.paste(part2, (0, 0, x_size - delta, y_size))
        im.save("roll_img.png", "PNG")
except IOError:
    pass
