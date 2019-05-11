import curses
from curses.textpad import Textbox, rectangle


def main(screen):
    screen.addstr(0, 0, "Enter text")

    edit_win = curses.newwin(5, 30, 2, 1)
    rectangle(screen, 1, 0, 1 + 5 + 1, 1 + 30 + 1)

    screen.refresh()

    box = Textbox(edit_win)
    box.edit()

    message = box.gather()
    return message


msg = curses.wrapper(main)
print(msg)
