import time
from math import sqrt
from joblib import Parallel, delayed

start1 = time.time()
lst = [sqrt(i ** 2) for i in range(10)]
end1 = time.time() - start1
print(lst, end1)

start2 = time.time()
lst = Parallel(n_jobs=2,prefer="threads")(delayed(sqrt)(i ** 2) for i in range(10))
end2 = time.time() - start2
print(lst, end2)

for i in range(10):
    print(sqrt(i ** 2))