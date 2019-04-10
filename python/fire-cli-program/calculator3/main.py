import fire


class Calculator(object):

    @classmethod
    def add(cls, x, y):
        return x + y

    @classmethod
    def sub(cls, x, y):
        return x - y


if __name__ == '__main__':
    fire.Fire(Calculator)
