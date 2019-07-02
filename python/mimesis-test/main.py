from mimesis import Person
from mimesis.enums import Gender

person = Person('en')

print(person.full_name(gender=Gender.FEMALE))
print(person.full_name(gender=Gender.MALE))

from mimesis import Generic

generic = Generic('ko')

print(generic.datetime.month())
print(generic.code.imei())
print(generic.food.fruit())

from mimesis import Address

address_ko = Address('ko')

print(address_ko.region())
print(address_ko.federal_subject())
print(address_ko.address())
print(address_ko.address())

from mimesis import Person

person = Person('ko', seed=0xffff)
print(person.full_name())

from mimesis.schema import Field, Schema
from mimesis.enums import Gender

field = Field('ko')
schema_dict = (lambda: {
    'version': field('version', pre_release=True),
    'timestamp': field('timestamp', posix=False),
    'token': field('token_hex'),
    'owner': {
        'id': field('uuid'),
        'name': field('full_name', gender=Gender.FEMALE),
        'email': field('person.email', key=str.lower)
    },
})
schema = Schema(schema=schema_dict)
print(schema.create(iterations=2))
