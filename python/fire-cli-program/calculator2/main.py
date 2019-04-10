import fire


def add(x, y):
    return x + y


def sub(x, y):
    return x - y


if __name__ == '__main__':
    fire.Fire({
        'add': add,
        'subtract': sub,
    })
