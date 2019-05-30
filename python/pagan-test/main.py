import pagan

input_name = 'blog'

img = pagan.Avatar(input_name, pagan.SHA512)

img.show()

out_path = 'output/'
filename = input_name

img.save(out_path, filename)