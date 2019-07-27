from wsgiref.simple_server import make_server
from pyramid.config import Configurator
from pyramid.httpexceptions import HTTPFound
from pyramid.response import Response
from pyramid.view import view_config


@view_config(route_name='home')
def home_view(request):
    return Response('<p>Visit <a href="/login">login page</a></p>')


@view_config(route_name='login')
def login_view(request):
    return HTTPFound(location="/")


if __name__ == '__main__':
    with Configurator() as config:
        config.add_route('home', '/')
        config.add_route('login', '/login')
        config.scan('view')
        app = config.make_wsgi_app()
    server = make_server('0.0.0.0', 8000, app)
    server.serve_forever()