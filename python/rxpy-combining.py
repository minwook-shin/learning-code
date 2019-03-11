from rx import *


class ObserverClass(Observer):

    def on_next(self, value):
        print("print : " + str(value))

    def on_completed(self):
        print("Done!")

    def on_error(self, error):
        print("Error : "+str(error))


fisrt = Observable.range(1, 5)
second = Observable.just(20)
fisrt.merge(second).subscribe(ObserverClass())

Observable.repeat(Observable.from_((1, 2, 3)),
                  3).merge_all().subscribe(ObserverClass())


ob1 = Observable.from_((1, 2))
ob2 = Observable.from_((3, 4))
Observable.concat([ob2, ob1]).subscribe(ObserverClass())


obRange = Observable.range(0, 6)
Observable.zip(obRange, obRange.skip(1), obRange.skip(2), lambda s1,
               s2, s3: '%s / %s / %s' % (s1, s2, s3)).subscribe(ObserverClass())


Observable.zip_list(obRange, obRange.skip(
    1), obRange.skip(2)).subscribe(ObserverClass())


s1 = Observable.interval(0).map(lambda i: 'First : %s' % i)
s2 = Observable.interval(0).map(lambda i: 'Second: %s' % i)
s1.combine_latest(s2, lambda s1, s2: '%s, %s'
                  % (s1, s2)).take(6).subscribe(ObserverClass())


s1 = Observable.interval(0).map(lambda i: 'First : %s' % i)
s2 = Observable.interval(0).map(lambda i: 'Second: %s' % i)
s1.with_latest_from(s2, lambda s1, s2: '%s, %s' %
                    (s1, s2)).take(6).subscribe(ObserverClass())

s1 = Observable.interval(0).map(lambda i: 'First : %s' % i)
s2 = Observable.interval(0).map(lambda i: 'Second: %s' % i)
s1.join(s2, lambda _: Observable.timer(10), lambda _: Observable.timer(
    0), lambda x, y: '%s %s' % (x, y)).take(6).subscribe(ObserverClass())

# .and_().then_do()


Observable.range(0, 3).select(lambda x: Observable.range(x, 3).map(
    lambda v: '%s (from stream nr %s)' % (v, x))).switch_latest().subscribe(ObserverClass())
