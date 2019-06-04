import joblib

filename = './test.joblib'

to_persist = [('text', [1, 2, 3])]

joblib.dump(to_persist, filename)

print(joblib.load(filename))


joblib.dump(to_persist, filename + '.gz', compress=('gzip', 3))
print(joblib.load(filename + '.gz'))


import gzip
with gzip.GzipFile(filename + '.gz', 'wb', compresslevel=3) as f:
    joblib.dump(to_persist, f)
with gzip.GzipFile(filename + '.gz', 'rb') as f:
    print(joblib.load(f))