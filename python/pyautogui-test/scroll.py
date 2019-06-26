import time

import pyautogui

time.sleep(3)

width, height = pyautogui.size()
pyautogui.moveTo(width / 2, height / 2)

for i in range(10):
    pyautogui.scroll(-10, x=width / 2, y=height / 2)
    time.sleep(0.5)
