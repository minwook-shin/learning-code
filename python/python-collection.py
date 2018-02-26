import collections as c

stack = []
print(stack)
stack.append(10)
print(stack)
stack.append(20)
print(stack)
stack.append(30)
print(stack)
stack.pop()
print(stack)
stack.pop()
print(stack)

queue = []
print(queue)
queue.append(10)
print(queue)
queue.append(20)
print(queue)
queue.append(30)
print(queue)
queue.pop(0)
print(queue)
queue.pop(0)
print(queue)

t = (10,20,30)
print(t)

s = set([10,20,30,30,20,10])
print(s)
s.add(10)
print(s)
s.remove(10)
print(s)
s.clear()
print(s)

d = {}
d = {"fisrt" : 10,"second" : 20, "third" : 30}
print(d)
print(d.items())
print(d.keys())
print(d.values())

deque_list = c.deque()
for i in range(5):
    deque_list.append(i)
print(deque_list)
for i in range(len(deque_list)):
    deque_list.rotate(i)
    print(deque_list)
print(c.deque(reversed(deque_list)))

o = c.OrderedDict()

o["1"] = 1
o["4"] = 4
o["3"] = 3
o["2"] = 2
for k,v in o.items():
    print(k,v)

d = c.defaultdict(lambda: 0) # Default 값을 0으로 설정합 
print(d["first"])

count = c.Counter("hello, world!")
print(count)

n = c.namedtuple("hello",["x","y"])
n2= n(11,y=12)
print(n2)