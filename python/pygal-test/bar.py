import pygal

bar_chart = pygal.Bar()
bar_chart.title = "bar"
bar_chart.add('test', [0, 1, 2, 3, 4, 5, 6, 7, 8])
bar_chart.render_to_file('bar.svg')