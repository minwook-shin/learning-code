from blinker import signal


def subscribe_function(sender):
    print(sender)


sig = signal('ready')
sig.connect(subscribe_function)


class Processor:
    def __init__(self, name):
        self.name = name

    def executor(self):
        ready = signal('ready')
        ready.send(self)
        print("Processing.")

    def __repr__(self):
        return '<Processor %s>' % self.name


if __name__ == '__main__':
    test = Processor('test')
    test.executor()