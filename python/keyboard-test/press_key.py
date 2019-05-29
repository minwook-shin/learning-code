import keyboard


def print_pressed_keys(e):
    for code in keyboard._pressed_events:
        line = ', '.join(str(code))
        print(line,e.name)


keyboard.hook(print_pressed_keys)
keyboard.wait('esc')