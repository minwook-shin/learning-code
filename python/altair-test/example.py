import altair as alt

from vega_datasets import data

print(data.cars())

cars = data.cars()

chart = alt.Chart(cars).mark_point().encode(
    x='Horsepower',
    y='Miles_per_Gallon',
    color='Origin',
)

chart.save('chart.png')