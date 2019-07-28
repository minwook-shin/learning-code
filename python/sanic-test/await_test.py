from sanic import Sanic
from sanic import response
import asyncio

app = Sanic(__name__)


@app.route("/await")
async def test_await(request):
    await asyncio.sleep(5)
    return response.text("async 5 second")


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=8000, debug=True)
