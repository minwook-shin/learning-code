import urwid

choices = ["first", "second", "third"]


def menu(title, choices):
    body = [urwid.Text(title), urwid.Divider()]
    for c in choices:
        button = urwid.Button(c)
        urwid.connect_signal(button, 'click', exit_program)
        body.append(urwid.AttrMap(button, None, focus_map='reversed'))
    return urwid.ListBox(urwid.SimpleFocusListWalker(body))


def exit_program(button):
    raise urwid.ExitMainLoop()


main = urwid.Padding(menu("title", choices))
top = urwid.Overlay(main, urwid.SolidFill(u'\N{MEDIUM SHADE}'),
                    align='center', width=('relative', 10),
                    valign='middle', height=('relative', 50))
urwid.MainLoop(top).run()