import pyautogui

pyautogui.FAILSAFE = True


width, height = pyautogui.size()
pyautogui.moveTo(width / 2, height / 2)
