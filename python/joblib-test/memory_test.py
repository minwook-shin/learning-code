
from joblib import Memory


cache_dir = './cache_dir'
memory = Memory(cache_dir, verbose=0)


@memory.cache
def func(x):
    print('Running func(%s)' % x)
    return x


print(func(1))

print(func(1))

print(func(2))