from werkzeug.wrappers import Request, Response
from werkzeug.serving import run_simple


@Request.application
def application(request):
    return Response('Hello, World!')


if __name__ == '__main__':
    run_simple('localhost', 8000, application)
