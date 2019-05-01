from dateutil.relativedelta import *
from dateutil.easter import *
from dateutil.rrule import *
from dateutil.parser import *

now = parse("Wed May 01 01:01:50 UTC 2019")
print(now)

today = now.date()
print(today)


count = list(rrule(DAILY, count=10, dtstart=now))
print(count)


year = rrule(YEARLY, dtstart=now, bymonth=6, bymonthday=23, byweekday=FR)[0].year
print(year)
relative_delta = relativedelta(easter(year), today)
print(relative_delta)
print(today + relative_delta)
