import os
from PIL import Image

input_img = input()
f, e = os.path.splitext(input_img)

try:
    Image.open(input_img).save("original.jpeg")
except IOError:
    pass
