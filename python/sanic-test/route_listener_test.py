from sanic import Sanic
from sanic.log import logger as log
from sanic import response

app = Sanic(__name__)


@app.route("/")
async def test_async(request):
    return response.json({"test": True})


@app.listener('before_server_start')
def before_start(app, loop):
    log.info("before_server_start")


@app.listener('after_server_start')
def after_start(app, loop):
    log.info("after_server_start")


@app.listener('before_server_stop')
def before_stop(app, loop):
    log.info("before_server_stop")


@app.listener('after_server_stop')
def after_stop(app, loop):
    log.info("after_server_stop")


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=8000, debug=True)
