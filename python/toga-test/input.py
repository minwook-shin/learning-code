import toga
from toga.style.pack import *


def build(app):
    def calculate(widget):
        try:
            result_form.value = int(input_form.value) * int(input_form.value)
        except ValueError:
            result_form.value = 'ValueError'

    box = toga.Box()

    input_box = toga.Box()
    box.add(input_box)
    input_form = toga.TextInput()
    input_form.style.update(flex=1)
    input_box.add(input_form)
    input_box.style.update(direction=ROW)

    result_box = toga.Box()
    box.add(result_box)
    result_form = toga.TextInput(readonly=True)
    result_box.add(result_form)
    result_box.style.update(direction=ROW)
    result_form.style.update(flex=1)

    button = toga.Button('confirm', on_press=calculate)
    button.style.update(flex=1)
    box.add(button)
    box.style.update(direction=COLUMN)

    return box


def main():
    return toga.App(name='Converter', app_id='io.github.minwook-shin.Converter', startup=build)


if __name__ == '__main__':
    main().main_loop()