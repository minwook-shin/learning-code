import gevent
from gevent import monkey

monkey.patch_all()

import requests

urls = [
    'https://www.google.com/',
    'https://www.python.org/'
]


def print_head(url):
    data = requests.get(url).text
    print('%s:  %r' % (url, data))


jobs = [gevent.spawn(print_head, _url) for _url in urls]
gevent.wait(jobs)