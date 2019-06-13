from marshmallow import Schema, fields, post_load, pprint


class User(object):
    def __init__(self, name, email):
        self.name = name
        self.email = email


class UserSchema(Schema):
    name = fields.Str()
    email = fields.Email()

    # @post_load
    # def make_user(self, data):
    #     return User(**data)


user_data = {
    'email': u'name@exmple.com',
    'name': u'name'
}

schema = UserSchema()
result = schema.load(user_data)
pprint(result.data)