from bidict import bidict, IGNORE

hello_world = bidict({'hello': 'world'})

symbol = hello_world['hello']
inverse_symbol = hello_world.inverse['world']

print(hello_world)

print(symbol)
print(inverse_symbol)

print(hello_world.inverse)

test_value = bidict({'First': 'one', 'Second': 'Two', 'Third': 'Three'})
test_value['First'] = 'One'

print(sorted(test_value.keys()))
print(sorted(test_value.values()))

print(test_value)
del test_value.inverse['Three']
print(test_value)

print('First' in test_value)

print(test_value.get('First', 'default'))
print(test_value.get('Fake_First', 'default'))

print(test_value.pop('Second', 'default'))
print(test_value)

test_value.update(new_key="new_value")
print(test_value)

# test_value.update(new_key2="new_value")
# print(test_value)
test_value.forceput('new_key2','new_value')
print(test_value)

test_value.putall([('new_key', 'new_value')], on_dup_val=IGNORE)
print(test_value)


print(bidict(hello='world') == dict(hello='world'))
