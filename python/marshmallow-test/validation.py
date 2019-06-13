from marshmallow import fields, Schema, validates, ValidationError, pprint


class UserSchema(Schema):
    name = fields.Str(required=True)
    email = fields.Email()


data, errors = UserSchema().load({'email': 'foo'})
print(data, errors)


class ItemSchema(Schema):
    quantity = fields.Integer()

    @validates('quantity')
    def validate_quantity(self, value):
        if value < 0:
            raise ValidationError('Quantity must be greater than 0.')
        if value > 10:
            raise ValidationError('Quantity must not be greater than 30.')


in_data = {'quantity': 11}
result, errors = ItemSchema().load(in_data)
print(result, errors)

