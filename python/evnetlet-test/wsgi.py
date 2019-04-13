import eventlet
from eventlet import wsgi


def main(_env, response):
    response('200 OK', [('Content-Type', 'text/plain')])
    return ['Hello, World!']


wsgi.server(eventlet.listen(('127.0.0.1', 8080)), main)