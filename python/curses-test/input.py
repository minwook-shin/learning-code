import curses


def main(screen):
    screen.clear()
    screen.addstr(0, 0, 'Hello World')
    while True:
        c = screen.getch()
        screen.addstr(1, 0, str(c))


curses.wrapper(main)
