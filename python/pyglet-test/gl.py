from pyglet.gl import *

win = pyglet.window.Window()


@win.event
def on_draw():
    glBegin(GL_LINES)
    glVertex3f(100.0, 100.0, 0)
    glVertex3f(400.0, 100.0, 0)
    glEnd()

    glBegin(GL_LINES)
    glVertex3f(100.0, 400.0, 0)
    glVertex3f(100.0, 100.0, 0)
    glEnd()

    glBegin(GL_LINES)
    glVertex3f(400.0, 100.0, 0)
    glVertex3f(400.0, 400.0, 0)
    glEnd()

    glBegin(GL_LINES)
    glVertex3f(100.0, 400.0, 0)
    glVertex3f(400.0, 400.0, 0)
    glEnd()


pyglet.app.run()
