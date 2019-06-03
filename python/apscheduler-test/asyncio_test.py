from datetime import datetime
from apscheduler.schedulers.asyncio import AsyncIOScheduler
import asyncio


def hello_world():
    print('hello world! time : %s' % datetime.now())


scheduler = AsyncIOScheduler()
scheduler.add_job(hello_world, 'interval', seconds=5)
scheduler.start()

try:
    asyncio.get_event_loop().run_forever()
except KeyboardInterrupt:
    pass
