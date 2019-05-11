from curses import wrapper


def main(screen):
    screen.clear()

    screen.addstr(0, 0, "hello,world!")

    screen.refresh()
    screen.getkey()


wrapper(main)
