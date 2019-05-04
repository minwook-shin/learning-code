import maya

import time
from datetime import datetime

now = maya.now()
print(now)

tomorrow = maya.when('tomorrow')
print(tomorrow)

ds = tomorrow.slang_date()
print(ds)

st = tomorrow.slang_time()
print(st)

iso = tomorrow.iso8601()
print(iso)

rtc2822 = tomorrow.rfc2822()
print(rtc2822)

rtc3339 = tomorrow.rfc3339()
print(rtc3339)

dt = tomorrow.datetime()
print(dt)
print(type(dt))

pd = maya.parse('2019-05-04 18:23:45.423992+00:00').datetime(to_timezone="Asia/Seoul", naive=True)
print(pd)

fd = maya.MayaDT.from_datetime(datetime.utcnow())
print(fd)

fs = maya.MayaDT.from_struct(time.gmtime())
print(fs)

md = maya.MayaDT(time.time())
print(md)

when = maya.when('2019-05-04', timezone="Asia/Seoul")
print(when)
print(when.day)
print(when.add(days=5).day)
print(when.timezone)

interval = maya.intervals(start=maya.now(), end=maya.now().add(days=1), interval=60 * 60)
print(interval)
for i in interval:
    print(i)

snap = maya.when('Mon, 06 May 2019 21:21:42 GMT').snap('@d+3h').rfc2822()
print(snap)
