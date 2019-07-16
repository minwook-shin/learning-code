import matplotlib.pyplot as plt
import numpy as np
from moviepy.editor import VideoClip
from moviepy.video.io.bindings import mplfig_to_npimage

x = np.linspace(-np.pi, np.pi)

fig, ax = plt.subplots()


def make_frame(t):
    ax.clear()
    ax.plot(x, np.sin(x + 2 * np.pi / 2 * t))
    return mplfig_to_npimage(fig)


animation = VideoClip(make_frame, duration=2)
animation.write_gif('matplotlib.gif', fps=60)
