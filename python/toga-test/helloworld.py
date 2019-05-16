import toga


def button_handler(widget):
    print("hello")


def build(app):
    box = toga.Box()
    button = toga.Button('Hello world', on_press=button_handler)
    box.add(button)
    return box


def main():
    return toga.App(name='First App', app_id='io.github.minwook-shin.helloworld', startup=build)


if __name__ == '__main__':
    main().main_loop()