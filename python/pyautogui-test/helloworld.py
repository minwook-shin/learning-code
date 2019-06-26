import pyautogui

pyautogui.FAILSAFE = True

width, height = pyautogui.size()
pyautogui.moveTo(width / 2, height / 2)

pyautogui.click()
pyautogui.doubleClick()
pyautogui.typewrite('print("Hello world!")', interval=0.2)
pyautogui.keyDown('shift')
pyautogui.press(['left', 'left', 'left', 'left', 'left', 'left', 'left', 'left'])
pyautogui.keyUp('shift')
