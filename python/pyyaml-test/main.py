import yaml

test = yaml.load("""
- first
- second
- third
""")

print(test)

yaml.warnings({'YAMLLoadWarning': False})

# or

test = yaml.load("""
- first
- second
- third
""", Loader=yaml.FullLoader)
print(test)

doc = """
---
name: hello world
description: >
    testing file
---
name: readme
description: >
    readme file
---
name: make
description: >
    build file
"""

for data in yaml.load_all(doc):
    print(data)

test = yaml.load("""
none: [~, null]
bool: [true, false, on, off]
int: 10
float: 3.14
list: [string,string]
dict: {id: "test_id", count: 5}
""")
print(test)


class TestClass:
    def __init__(self, id_, name, count):
        self.id = id_
        self.name = name
        self.count = count


obj = yaml.load("""
!!python/object:__main__.TestClass
id: test_id
name: test_name
count: 1
""")
print(obj.id)
print(obj.name)
print(obj.count)

print(yaml.dump({'name': 'test_name', 'email': 'example@example.com', 'todo': ['post article', 'sql study']}))

print(yaml.dump(TestClass(id_="test_id", name="test_name", count=1)))
