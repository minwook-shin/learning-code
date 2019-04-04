from pymongo import MongoClient

client = MongoClient('localhost', 27017)
client.drop_database('test_database')
db = client.test_database
collection = db.test_collection

user = {"name": "Kim",
        "birthday": 0000}

users = collection.users
user_id = users.insert_one(user).inserted_id

print(user)
print(users.find_one())
print(users.find_one({"name": "Kim"}))
print(users.find_one({"_id": user_id}))

user_collection = [{"name": "Lee",
                    "birthday": 1111},
                   {"name": "Shin",
                    "birthday": 2222}]

user_ids = users.insert_many(user_collection)

for i in users.find():
    print(i)
