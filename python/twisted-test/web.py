from twisted.web import server, resource
from twisted.internet import reactor, endpoints


class WebServer(resource.Resource):
    isLeaf = True

    def render_GET(self, request):
        request.setHeader(b"content-type", b"text/plain")
        content = u"this is web server."
        return content.encode("ascii")


endpoints.serverFromString(reactor, "tcp:8080").listen(server.Site(WebServer()))
reactor.run()
