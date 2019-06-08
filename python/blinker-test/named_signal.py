from blinker import signal


named = signal('named')
print(named is signal('named'))