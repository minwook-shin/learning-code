from asciimatics.screen import Screen


def main(screen):
    while True:
        screen.print_at('input : ', 0, 0)
        screen.refresh()
        ev = screen.get_key()
        if ev in (ord('Q'), ord('q')):
            return
        if ev is not None:
            screen.print_at(chr(ev), 0, 1)
        screen.refresh()


Screen.wrapper(main)
