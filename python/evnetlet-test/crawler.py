import eventlet
from eventlet.green.urllib import request


urls = [
    "https://www.google.com/intl/en_ALL/images/logo.gif",
    "http://python.org/images/python-logo.gif",
]


def fetch(text):
    content = request.urlopen(text).read()
    print("done : ", text)
    return text, content


pool = eventlet.GreenPool(200)
for url, body in pool.imap(fetch, urls):
    print("got body from", url)