from asciimatics.screen import Screen
from time import sleep


def main(screen):
    screen.print_at('Hello world!', 0, 0)
    screen.print_at('Hello world!', 0, 1)
    screen.refresh()
    screen.move(0, 2)
    screen.draw(10, 2)
    screen.refresh()
    sleep(10)


if __name__ == '__main__':
    Screen.wrapper(main)