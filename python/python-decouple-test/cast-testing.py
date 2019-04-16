from decouple import config


debug = config('DEBUG', cast=bool)
print(debug)
print(type(debug))

db_url = config('DATABASE_URL', cast=str)
print(db_url)
print(type(db_url))

host = config('HOST', cast=str, default='localhost')
print(host)
print(type(host))

port = config('PORT', cast=int, default = "8080")
print(port)
print(type(port))

