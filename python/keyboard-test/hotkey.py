import keyboard

shortcut = keyboard.read_hotkey()


def on_triggered():
    print("Triggered!")


keyboard.add_hotkey(shortcut, on_triggered)
keyboard.wait('esc')
