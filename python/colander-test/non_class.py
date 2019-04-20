import colander

schema = colander.SchemaNode(colander.Mapping())
schema.add(colander.SchemaNode(colander.String(), name='id'))
schema.add(colander.SchemaNode(colander.Int(), name='login',
                               validator=colander.Range(0, 999)))

my_information = {
    'id': 'example@example.com',
    'login': '1',
}

try:
    deserialized = schema.deserialize(my_information)
except colander.Invalid as e:
    print(e)

serialized = schema.serialize(my_information)
print(serialized)
