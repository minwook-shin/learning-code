import rapidjson

with open("sample.json", "rb") as f:
    sample = rapidjson.loads(f.read())

    print(sample)
    print(type(sample))


print(rapidjson.dumps(sample))
print(type(rapidjson.dumps(sample)))


class Stream:

    def write(self, data):
        print("Chunk:", data)


rapidjson.dump(sample, Stream(), chunk_size=100)