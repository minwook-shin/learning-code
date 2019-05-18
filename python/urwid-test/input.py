import urwid


def exit_program(key):
    if key in ('q', 'Q'):
        raise urwid.ExitMainLoop()
    txt.set_text(repr(key))


txt = urwid.Text("")
fill = urwid.Filler(txt, 'top')
loop = urwid.MainLoop(fill, unhandled_input=exit_program)
loop.run()