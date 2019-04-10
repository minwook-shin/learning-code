import fire


class Main(object):

    def __init__(self, text):
        self.text = text


if __name__ == '__main__':
    fire.Fire(Main)
