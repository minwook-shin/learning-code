from datetime import datetime
from apscheduler.schedulers.blocking import BlockingScheduler


def hello_world():
    print('hello, world! time : %s' % datetime.now())


scheduler = BlockingScheduler()
scheduler.add_job(hello_world, 'interval', seconds=5)

try:
    scheduler.start()
except KeyboardInterrupt:
    pass