from schematics.models import Model
from schematics.transforms import blacklist
from schematics.types import StringType, URLType
from schematics.exceptions import DataError


import json


class TestClass(Model):
    uuid = StringType(required=True)
    hello = StringType(required=True)
    url = URLType()

    class Options:
        roles = {'public': blacklist('uuid')}
        serialize_when_none = False


mock1 = TestClass.get_mock_object().to_primitive()
print(mock1)

mock2 = TestClass.get_mock_object().to_primitive(role='public')
print(mock2)

example1 = TestClass({'hello': 'world', 'url': 'https://www.python.org', 'uuid': '0000'})


json = json.dumps(example1.to_primitive())
print(json)


example2 = TestClass()
example2.uuid = '0000'
example2.hello = 'world'
example2.website = 'https://www.python.org'

print(example2.uuid)
print(example2.hello)
print(example2.website)


try:
    example2.validate()
except DataError as e:
    print(e)