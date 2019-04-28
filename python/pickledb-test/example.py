import pickledb


db = pickledb.load('test.db', False)

db.set('key', 'value')


value = db.get('key')
print(value)

value_all = db.getall()
print(value_all)

db.set('key1', 'temp_value')
value_all = db.getall()
print(value_all)
db.rem('key1')
value_all = db.getall()
print(value_all)

total = db.totalkeys()
print(total)


db.append('key','_test')
value_all = db.get('key')
print(value_all)


is_key = db.exists('key')
print(is_key)

db.lcreate("test_list")
db.ladd('test_list','val')
print(db.lget('test_list',0))

is_dump = db.dump()
print(is_dump)
