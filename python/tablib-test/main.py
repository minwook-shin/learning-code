import tablib

headers = ('id', 'name')

data = [
    ('test_id_1', 'name_1'),
    ('test_id_2', 'name_2')
]

data = tablib.Dataset(*data, headers=headers)
print(data)
print(type(data))

data.append(('test_id_3', 'name_3'))
data.append_col((1, 2, 3), header='count')

print(data)
print(data['id'])