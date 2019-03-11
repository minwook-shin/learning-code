from rx import *
import random
import time


class ObserverClass(Observer):

    def on_next(self, value):
        print("print : " + str(value))

    def on_completed(self):
        print("Done!")

    def on_error(self, error):
        print("Error : "+str(error))


Observable.just({'answer': random.randint(0, 10)}).subscribe(ObserverClass())
time.sleep(0.5)
Observable.just({'answer': random.randint(0, 10)}).map(
    lambda x: x.get('answer') * 2).subscribe(ObserverClass())

Observable.start(func=lambda: random.randint(0, 10)).subscribe(ObserverClass())

Observable.range(0, 3).subscribe(ObserverClass())

Observable.from_iterable((1, 2, random.randint(0, 10))
                         ).subscribe(ObserverClass())

gen = (random.randint(0, 10) for j in range(3))
Observable.from_iterable(gen).subscribe(ObserverClass())


def g(f, a, b):
    f(a, b)
    print('called f')


Observable.from_callback(lambda a, b, f: g(f, a, b))(
    'fu', 'bar').subscribe(ObserverClass())


Observable.repeat({'rand': time.time()}, 3).subscribe(ObserverClass())

Observable.generate(0, lambda x: x < 5, lambda x:  x + 1.0,
                    lambda x: x * 2).subscribe(ObserverClass())


Observable.defer(lambda: Observable.just(
    random.randint(0, 10))).subscribe(ObserverClass())

Observable.if_then(
    lambda: True, Observable.return_value(43), Observable.return_value(56)).subscribe(ObserverClass())

Observable.empty().subscribe(ObserverClass())

Observable.never().subscribe(ObserverClass())

Observable.throw(ZeroDivisionError).subscribe(ObserverClass())
