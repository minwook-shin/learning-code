import keyboard
import fileinput

for line in fileinput.input():
	keyboard.write(line)