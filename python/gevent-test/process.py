import gevent
from gevent import subprocess


p1 = subprocess.Popen(['uname'], stdout=subprocess.PIPE)
p2 = subprocess.Popen(['ls'], stdout=subprocess.PIPE)

gevent.wait([p1, p2], timeout=2)

if p1.poll() is not None:
    print('%r' % p1.stdout.read())
if p2.poll() is not None:
    print('%r' % p2.stdout.read())

p1.stdout.close()
p2.stdout.close()