import time
import gevent
from gevent.threadpool import ThreadPool


pool = ThreadPool(3)
for _ in range(4):
    pool.spawn(time.sleep, 1)
gevent.wait()
