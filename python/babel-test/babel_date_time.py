from datetime import date
from babel.dates import format_date

d = date(2019, 6, 1)
print(format_date(d, locale='ko'))
print(format_date(d, locale='ko_KR'))

print(format_date(d, format='short', locale='ko'))
print(format_date(d, format='long', locale='ko'))
print(format_date(d, format='full', locale='ko'))


from datetime import timedelta
from babel.dates import format_timedelta
delta = timedelta(days=6)
print(format_timedelta(delta, locale='ko_KR'))