from moviepy.editor import *
from moviepy.video.tools.segmenting import findObjects

text_clip = TextClip('Hello, World!', color='white', kerning=5, fontsize=100)
composite_video_clip = CompositeVideoClip([text_clip.set_pos('center')], size=(1000, 1000))


def in_effect(pos):
    d = lambda t: 1.0 / (0.3 + t ** 8)
    return lambda t: pos + 400 * d(t)


def out_effect(pos):
    d = lambda t: max(0, t)
    return lambda t: pos + 400 * d(t)


def move(letters, func_pos):
    return [i.set_pos(func_pos(i.screenpos)) for i in letters]


clips = [CompositeVideoClip(move(findObjects(composite_video_clip), func_pos), size=(1000, 1000)).subclip(0, 5)
         for func_pos in [in_effect, out_effect]]

video_clip = concatenate_videoclips(clips)
video_clip.write_videofile('hello_world.avi', fps=60, codec='mpeg4')
