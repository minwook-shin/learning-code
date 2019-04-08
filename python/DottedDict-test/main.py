from dotted.collection import DottedCollection, DottedDict, DottedList
from dotted.utils import dot, dot_json

dotted_arr = DottedList([[0, 1], [2, 3], [4, 5]])

print(dotted_arr[0])
print(dotted_arr['0.0'])

print(dotted_arr['1'])
print(dotted_arr['1.1'])

print(dotted_arr[2])

dotted_arr.append(11)
print(dotted_arr)

dotted_arr[len(dotted_arr)] = 12
print(dotted_arr)


dotted_dict = DottedDict({'hello': {'world': {'python': '3'}}})
print(dotted_dict['hello'])
print(dotted_dict['hello.world'])
print(dotted_dict['hello.world.python'])

print(dotted_dict.hello)
print(dotted_dict.hello.world)
print(dotted_dict.hello.world.python)

dotted_dict2 = DottedCollection.factory({
    'hello': [{'world': {'python': ['3', '7', '3']}}]
})
print(dotted_dict2['hello'][0]['world']['python'][0])
print(dotted_dict2.hello[0].world.python[0])
print(dotted_dict2.hello[0].world['python'][0])
print(dotted_dict2.hello[0].world['python.0'])
print(dotted_dict2.hello['0.world'].python[0])
print(dotted_dict2['hello.0.world.python.0'])

dotted_dict2['c++.path'] = {'hello': 'world'}
dotted_dict2['java.path'] = ['hello']
print(dotted_dict2)

dot_obj = dot({'hello': 'world'})
print(dot_obj)
dotted_json = dot_json('{"hello": "world"}')
print(dotted_json)
