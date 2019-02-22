import rx
from rx import Observable, Observer

# Using Subject and Stream
from rx.subjects import Subject

# Generating a sequence

print("Generating a sequence")


class ObserverClass(Observer):
    def on_next(self, value):
        print(value)

    def on_error(self):
        pass

    def on_completed(self):
        pass


src = Observable.from_iterable(range(5))
src.subscribe(ObserverClass())

print("--------")

src = Observable.from_iterable(range(5))
src.subscribe(print)

print("--------")

src = Observable.from_iterable(range(5))
src.subscribe(lambda x: print(x))

print("--------")

# Filtering a sequence

print("Filtering a sequence")


def func(x):
    return x % 2


src = Observable.from_(range(5))
src.filter(lambda x: func(x)).subscribe(print)

print("--------")

# Transforming a sequence

print("Transforming a sequence")


def func(x, i):
    print(i, ":", end=' ')
    return x * 2


src = Observable.from_(range(5))
src.map(func).subscribe(print)

print("--------")

# Merge

print("Merge")

src1 = Observable.range(1, 5)
src2 = Observable.from_("hello")
src1.merge(src2).subscribe(print)

print("--------")

## Subjects and Streams

print("Subjects and Streams")

ob = Subject()
ob.on_next(1)
sub1 = ob.subscribe(print)
ob.on_next(2)
sub2 = ob.subscribe(print)
sub1.dispose()
ob.on_next(3)
sub2.dispose()
