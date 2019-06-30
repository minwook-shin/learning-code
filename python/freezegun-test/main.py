from freezegun import freeze_time
import datetime


@freeze_time("2017-12-22")
def test():
    return datetime.datetime.now()


print(test())


@freeze_time("2017-12-22")
class Tester(object):
    def test_the_class(self):
        return datetime.datetime.now()


print(Tester().test_the_class())


@freeze_time("Jan 1th, 2018")
def test_human():
    return datetime.datetime.now()


test_human()

print(datetime.datetime.now())
with freeze_time("2017-12-22"):
    print(datetime.datetime.now())
print(datetime.datetime.now())

freezer = freeze_time("2017-12-22 12:00:01")
freezer.start()
print(datetime.datetime.now())
freezer.stop()


initial_datetime = datetime.datetime(year=1, month=1, day=1,
                                     hour=1, minute=1, second=1)
other_datetime = datetime.datetime(year=2, month=2, day=2,
                                   hour=2, minute=2, second=2)
with freeze_time(initial_datetime) as frozen_datetime:
    print(frozen_datetime())
    frozen_datetime.move_to(other_datetime)
    print(frozen_datetime())


@freeze_time("2017-12-22", as_arg=True)
def test(frozen_time):
    print(datetime.datetime.now())
    frozen_time.move_to("2018-01-01")
    print(datetime.datetime.now())


test()
