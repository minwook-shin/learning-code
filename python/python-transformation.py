from rx import *


class ObserverClass(Observer):

    def on_next(self, value):
        print("print : " + str(value))

    def on_completed(self):
        print("Done!")

    def on_error(self, error):
        print("Error : "+str(error))


Observable.from_((1, 2, 3)).map(lambda x: x * 2).subscribe(ObserverClass())

Observable.from_(range(10, 20, 2)).map(
    lambda x, i: "%s / %s" % (i, x * 2)).subscribe(ObserverClass())

Observable.from_([{'x': 1, 'y': 2}, {'x': 3, 'y': 4}]
                 ).pluck('y').subscribe(ObserverClass())


class test:
    def __init__(self, x, y):
        self.x = x
        self.y = y


Observable.from_([test(1, 2), test(3, 4)]).pluck_attr(
    'x').subscribe(ObserverClass())


Observable.range(1, 2).flat_map(
    lambda x: Observable.range(x, 2)).subscribe(ObserverClass())

Observable.range(1, 2).flat_map_latest(
    lambda x: Observable.range(x, 2)).subscribe(ObserverClass())


def res_sel(*a):
    return '-'.join([str(s) for s in a])


Observable.from_(('a', 'b', 'c')).flat_map(lambda x, i: (
    x, i), res_sel).subscribe(ObserverClass())

Observable.for_in([1, 2, 3],
                  lambda i: Observable.from_(('a', 'b', 'c')).to_blocking().map(
    lambda letter: '%s%s' % (letter, i))).subscribe(ObserverClass())


Observable.from_(('a', 'b', 'c')).to_blocking(
).many_select(lambda x: x.first()).merge_all().subscribe(ObserverClass())

Observable.from_(('1', '2', '3')).to_blocking().scan(lambda x,
                                                     y: int(x) + int(y)).subscribe(ObserverClass())


Observable.from_(('1', '2', '3')).to_blocking(
).timestamp().pluck_attr('timestamp').subscribe(ObserverClass())

Observable.from_(('1', '2', '3')).to_blocking().time_interval().map(
    lambda x: x.interval).subscribe(ObserverClass())
