from bokeh.plotting import figure, output_file, show

x = [1, 2, 3, 4, 5]
y = [1, 2, 3, 4, 5]

output_file("hello.html")

p = figure(title="simple line example", x_axis_label='x', y_axis_label='y')

p.line(x, y)

show(p)