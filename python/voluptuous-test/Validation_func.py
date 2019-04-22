from datetime import datetime
from voluptuous import Schema


def check_date(fmt='%Y-%m-%d'):
    return lambda v: datetime.strptime(v, fmt)


schema = Schema(check_date())
schema('2019-04-23')