from apscheduler.schedulers.blocking import BlockingScheduler

scheduler = BlockingScheduler()
scheduler.add_job('sys:stdout.write', 'interval', seconds=5, args=['hello,world!\n'])

try:
    scheduler.start()
except KeyboardInterrupt:
    pass