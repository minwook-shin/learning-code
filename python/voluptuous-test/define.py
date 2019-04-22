from voluptuous import Schema
from voluptuous import Url


schema = Schema(10)
num = schema(11)
print(num)

schema = Schema('string')
string = schema('string')
print(string)

schema = Schema(int)
int_type = schema(1)
print(int_type)

schema = Schema(Url())
url_type = schema('http://python.org')
print(url_type)

schema = Schema([10, 'a', 'string'])
ex1 = schema([10])
ex2 = schema([10, 'a'])
print(ex1)
print(ex2)

schema = Schema({1})
set = schema({1})
print(set)

schema = Schema({'hello': 'world'})
dict_type = schema({'hello': 'world'})
print(dict_type)