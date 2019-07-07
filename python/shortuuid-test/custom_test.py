import shortuuid

print(shortuuid.uuid())

print(shortuuid.uuid(name="example.com"))
print(shortuuid.uuid(name="http://www.google.com"))

print(shortuuid.ShortUUID().random(length=50))

print(shortuuid.get_alphabet())

import uuid

s = shortuuid.encode(uuid.uuid4())
print(s)


su = shortuuid.ShortUUID()
print(su.uuid())