import simdjson

with open('sample.json', 'rb') as f:
    document = simdjson.loads(f.read())

print(document)
print(type(document))

print(document["type"])
print(document["created_at"])
print(document["id"])

print(document["actor"])
for k, v in document["actor"].items():
    print(k, v)

print(document["repo"])
for k, v in document["repo"].items():
    print(k, v)

print(document["public"])

print(document["payload"])
for k, v in document["payload"].items():
    print(k, v)
