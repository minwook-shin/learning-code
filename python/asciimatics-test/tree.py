from asciimatics.effects import Print
from asciimatics.renderers import StaticRenderer
from asciimatics.scene import Scene
from asciimatics.screen import Screen

tree = r"""
       ${3,1}*
      / \
     /   \
    /_   _\
     /   \
    /     \
   /       \
  /__     __\
    /     \
   /       \
  /         \
 /___________\
      ${3}|||
      ${3}|||
"""


def main(screen):
    effects = [
        Print(screen, StaticRenderer(images=[tree]),
              colour=Screen.COLOUR_GREEN, x=0, y=0)
    ]
    screen.play([Scene(effects, -1)], stop_on_resize=True)


while True:
    Screen.wrapper(main)