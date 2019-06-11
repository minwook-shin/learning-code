from flask import request, url_for
from flask_api import FlaskAPI, status

app = FlaskAPI(__name__)

url_content = {
    0: 'main',
    1: 'shop',
    2: 'profile',
}


def content_repr(key):
    return {
        'url': request.host_url.rstrip('/') + url_for('url_detail', key=key),
        'text': url_content[key]
    }


@app.route("/", methods=['GET'])
def url_list():
    return [content_repr(idx) for idx in sorted(url_content.keys())]


@app.route("/<int:key>/", methods=['GET', 'PUT', 'DELETE'])
def url_detail(key):
    if request.method == 'PUT':
        note = str(request.data.get('text', ''))
        url_content[key] = note
        return content_repr(key)

    elif request.method == 'DELETE':
        url_content.pop(key, None)
        return '', status.HTTP_204_NO_CONTENT

    return content_repr(key)


if __name__ == "__main__":
    app.run(debug=True)
