from voluptuous import Schema, Required, All, Length, Range

schema = Schema({
    Required('id'): All(str, Length(min=1)),
    Required('login', default=5): All(int, Range(min=1, max=20)),
})

schema({'id': 'real_id', 'login': 10})