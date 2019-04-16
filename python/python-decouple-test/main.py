from decouple import config

SECRET_KEY = config('SECRET_KEY')
print(SECRET_KEY)

DEBUG = config('DEBUG', default=False, cast=bool)
print(DEBUG)

F_HOST = config('FAKE_HOST', default='localhost')
print(F_HOST)

F_PORT = config('FAKE_PORT', default=8080, cast=int)
print(F_PORT)
