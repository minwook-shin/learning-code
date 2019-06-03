import time
from datetime import datetime
from apscheduler.schedulers.background import BackgroundScheduler


def hello_world():
    print('hello, world! time : %s' % datetime.now())


scheduler = BackgroundScheduler()
scheduler.add_job(hello_world, 'interval', seconds=5)

scheduler.start()
time.sleep(10)


