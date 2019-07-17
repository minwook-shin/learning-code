from werkzeug.test import Client
from werkzeug.wrappers import BaseResponse

import hello_world

c = Client(hello_world.application, BaseResponse)
resp = c.get('/')
status = resp.status_code
print(status)
header = resp.headers
print(header)
response_data = resp.data
print(response_data)