from sanic import Sanic
from sanic import response
from sanic.exceptions import ServerError

app = Sanic(__name__)


@app.route("/exception")
def exception(request):
    raise ServerError("exception!")


@app.exception(ServerError)
async def test(request, exception):
    return response.json({"exception": "{}".format(exception)}, status=exception.status_code)


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=8000, debug=True)
