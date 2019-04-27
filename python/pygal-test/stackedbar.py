import pygal

bar_chart = pygal.StackedBar()
bar_chart.title = "Stacked Bar"
bar_chart.add('test', [0, 1, 2, 3, 4, 5, 6, 7, 8])
bar_chart.render_to_file('Stacked_Bar.svg')