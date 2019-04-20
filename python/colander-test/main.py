import colander


class Place(colander.MappingSchema):
    location = colander.SchemaNode(colander.String(),
                                   validator=colander.OneOf(['home', 'work']))
    Latitude_longitude = colander.SchemaNode(colander.String())


class GPS(colander.SequenceSchema):
    place = Place()


class Information(colander.MappingSchema):
    id = colander.SchemaNode(colander.String())
    login = colander.SchemaNode(
        colander.Int(), validator=colander.Range(0, 999))
    gps = GPS()


my_information = {
    'id': 'example@example.com',
    'login': '1',
    'gps': [{'location': 'home', 'Latitude_longitude': '37.576108, 126.976838'},
            {'location': 'work', 'Latitude_longitude': '37.400527, 127.104773'}],
}

schema = Information()

try:
    deserialized = schema.deserialize(my_information)
except colander.Invalid as e:
    print(e)

serialized = schema.serialize(my_information)
print(serialized)
