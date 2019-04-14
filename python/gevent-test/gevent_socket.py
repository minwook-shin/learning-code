import gevent
from gevent import socket


urls = ['www.google.com', 'www.python.org']

jobs = [gevent.spawn(socket.gethostbyname, url) for url in urls]
gevent.joinall(jobs, timeout=2)
content = [job.value for job in jobs]
print(content)