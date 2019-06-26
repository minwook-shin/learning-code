import pyautogui

pyautogui.FAILSAFE = True

width, height = pyautogui.size()

pyautogui.moveTo(width / 2, height / 2)
pyautogui.moveTo(500, 500, duration=2, tween=pyautogui.easeInOutBounce)

pyautogui.moveTo(width / 2, height / 2)
pyautogui.moveTo(500, 500, duration=2, tween=pyautogui.easeInOutQuad)

pyautogui.moveTo(width / 2, height / 2)
pyautogui.moveTo(500, 500, duration=2, tween=pyautogui.easeInOutBack)