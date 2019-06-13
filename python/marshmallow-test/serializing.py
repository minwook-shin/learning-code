from marshmallow import Schema, fields, pprint


class User(object):
    def __init__(self, name, email):
        self.name = name
        self.email = email


class UserSchema(Schema):
    name = fields.Str()
    email = fields.Email()


user = User(name="name", email="name@example.org")
schema = UserSchema()
result = schema.dump(user)
pprint(result.data)

result = schema.dumps(user)
pprint(result.data)
