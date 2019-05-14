import pyglet

window = pyglet.window.Window()
image = pyglet.image.load('image.png')
music = pyglet.resource.media('music.mp3')


@window.event
def on_draw():
    image.blit(0, 0)
    music.play()


pyglet.app.run()
