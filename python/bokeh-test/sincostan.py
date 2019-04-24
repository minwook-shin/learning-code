import numpy as np
from bokeh.layouts import gridplot
from bokeh.plotting import figure, output_file, show

x = np.linspace(0, 4*np.pi, 100)
y0 = np.sin(x)
y1 = np.cos(x)
y2 = np.tan(x)

output_file("lines.html")

s1 = figure()
s1.circle(x, y0)

s2 = figure()
s2.triangle(x, y1)

s3 = figure()
s3.square(x, y2)

p = gridplot([[s1, s2, s3]])

show(p)