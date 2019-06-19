import pyexcel

test_dict = [
    {
        "type": "PushEvent",
        "created_at": "2019-01-01T07:58:30Z",
        "id": "0000000001",
        "public": True,
    },
    {
        "type": "PushEvent",
        "created_at": "2019-01-02T07:58:30Z",
        "id": "0000000002",
        "public": True,
    },
    {
        "type": "PushEvent",
        "created_at": "2019-01-03T07:58:30Z",
        "id": "0000000003",
        "public": True,
    },
]
pyexcel.save_as(records=test_dict, dest_file_name="test.xls")

records = pyexcel.iget_records(file_name="test.xls")
for record in records:
    print(record['created_at'], record['id'])
