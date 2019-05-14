import pyglet
from pyglet.window import mouse, key

window = pyglet.window.Window()


@window.event
def on_key_press(symbol, modifiers):
    if symbol == key.A:
        print('A')
    elif symbol == key.S:
        print('S')
    elif symbol == key.D:
        print('D')
    elif symbol == key.F:
        print('F')

   elif symbol == key.LEFT:
        print('left')
    elif symbol == key.RIGHT:
        print('right')
    elif symbol == key.UP:
        print('up')
    elif symbol == key.DOWN:
        print('down')

@window.event
def on_mouse_press(x, y, button, modifiers):
    if button == mouse.LEFT:
        print('mouse left')
    elif button == mouse.RIGHT:
        print('mouse right')

@window.event
def on_draw():
    window.clear()

pyglet.app.run()
