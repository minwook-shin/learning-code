from decouple import config, Csv

example = config('INTEGERS', cast=Csv(int), default="1,2,3,4,5")
print(example)
print(type(example))

for i in example:
    print(i,end="\n")