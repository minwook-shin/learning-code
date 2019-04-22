from voluptuous import Schema

schema = Schema({
    'hello': str,
})


schema({'hello': 'world'})
