import fire


class Car1(object):

    @classmethod
    def run(cls):
        return 'running car1'


class Car2(object):

    @classmethod
    def run(cls):
        return "running car2"


class Cli(object):

    def __init__(self):
        self.car1 = Car1()
        self.car2 = Car2()

    def run(self):
        self.car1.run()
        self.car2.run()


if __name__ == '__main__':
    fire.Fire(Cli)
