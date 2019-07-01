import json

import httpretty
import urllib.request

@httpretty.activate
def test_http():
    httpretty.register_uri(
        httpretty.GET,
        "https://example.com",
        body='{"origin": "hello,world"}'
    )

    with urllib.request.urlopen('https://example.com') as f:
        response = f.read()

    print(json.loads(response.decode()))
    content = json.loads(response.decode())
    print(content == {'origin': 'hello,world'})


if __name__ == '__main__':
    test_http()