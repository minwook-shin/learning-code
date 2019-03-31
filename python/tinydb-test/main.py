from tinydb import TinyDB, Query, where

# from tinydb.storages import MemoryStorage
# db = TinyDB(storage=MemoryStorage)

db = TinyDB('db.json')
# table = db.table('table_name')

db.purge()

db.insert({'name': 'Kim', 'score': 90})
db.insert({'name': 'Shin', 'score': 80})
db.insert({'wrong-value': True})

db.insert_multiple([
    {'name': 'Lee', 'score': 80},
    {'name': 'Han', 'score': 50}])

db.upsert({'name': 'Han', 'score': 100}, Query().name == 'Han')

print(db.all())

for i in db:
    print(i)

print("name", db.search(Query().name == 'Shin'))
print("score", db.search(Query().score >= 90))
print("array indexing", db.search(Query()["name"] == 'Shin'))

print("traditional", db.search(where('name') == 'Shin'))

db.remove(~(Query().name.exists()))
print(db.all())

db.update({'score': 0}, Query().name == 'Han')
print(db.all())

shin_account = db.get(Query().name == 'Shin')
print(shin_account)

docs = db.search(Query().name == 'Shin')
for doc in docs:
    doc['score'] = 100
db.write_back(docs)
docs = db.search(Query().name == 'Shin')
print(docs)

print(len(db))
