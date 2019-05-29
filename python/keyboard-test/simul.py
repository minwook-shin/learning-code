import keyboard
import time


for i in range(20):
    keyboard.press('a')
    time.sleep(0.1)

keyboard.release('a')