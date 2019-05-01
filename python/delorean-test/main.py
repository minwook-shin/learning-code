from delorean import Delorean
from delorean import parse
from datetime import timedelta


d = Delorean()
print(d)

ds = d.shift("Asia/Seoul")
print(ds)

print(ds.next_sunday())

print(ds.replace(hour=12))

print(ds.truncate('day'))

print(ds.date)

print(ds.datetime)
print(type(ds.datetime))

print(ds.naive)

print(ds.epoch)

d = Delorean()
d + timedelta(hours=2)
d - timedelta(hours=2)


dp = parse("2019/05/02 00:00:00 -0700")
print(dp.date)