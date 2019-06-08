from blinker import signal


started = signal('started')


def round_function(num):
    print("Round : %s" % num)


def round_two(num):
    print(num)


started.connect(round_function)
started.connect(round_two, sender=2)

for i in range(1, 4):
    started.send(i)