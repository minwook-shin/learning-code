import tablib

headers = ('id', 'name')

data = [
    ('test_id_1', 'name_1'),
    ('test_id_2', 'name_2'),
    ('test_id_3', 'name_3')
]

data = tablib.Dataset(*data, headers=headers)
with open('test.xls', 'wb') as f:
    f.write(data.export('xls'))