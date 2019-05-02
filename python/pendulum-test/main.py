import pendulum

dt = pendulum.datetime(2019, 5, 1)
print(type(dt))

pendulum.datetime(2019, 5, 1, tz="Asia/Seoul")

dt = pendulum.local(2019, 5, 1)
print(dt.timezone.name)

now = pendulum.now()
print(now)

naive = pendulum.naive(2019, 5, 1)
print(naive.timezone)

dt = pendulum.parse('2019-05-01T22:00:00')
print(dt)

pendulum.set_locale('ko')
print(pendulum.now().add(years=1).diff_for_humans())

dt = pendulum.parse('2019-05-01T22:00:00')

print(dt.year)
print(dt.month)
print(dt.day)
print(dt.hour)
print(dt.minute)
print(dt.second)
print(dt.microsecond)
print(dt.day_of_week)
print(dt.day_of_year)
print(dt.week_of_month)
print(dt.week_of_year)
print(dt.days_in_month)

dt = pendulum.now()
print(dt.set(year=2017, month=12, day=22).to_datetime_string())

pendulum.now()
print(dt.to_date_string())

dt = pendulum.datetime(2019, 5, 1, 00, 00, 00)
print(dt.format('YYYY-MM-DD HH:mm:ss'))

dt = pendulum.now()
print(dt.format('[today] dddd'))

dt = pendulum.datetime(2019, 5, 1, 1, 1, 1)

print(dt.start_of('day'))