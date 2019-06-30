from doublex import Stub, ANY_ARG


class Collaborator:
    def except_method(self):
        pass

    def add_method(self, a, b):
        pass


with Stub(Collaborator) as stub:
    stub.except_method().raises(KeyError)
    stub.add_method(ANY_ARG).returns(4)

print(stub.add_method(2, 3))

try:
    stub.except_method()
except KeyError:
    print("key error!")



from doublex import Stub

with Stub() as stub:
    stub.example('test').returns(10)

# when
result = stub.example('test')
print(result)

from doublex import Spy


class Sender:
    def send_method(self, address, force=True):
        pass


sender = Spy(Sender)
sender.send_method("example@example.net")

try:
    sender.send_method()
except TypeError:
    print("argument error")

from doublex import Mock

with Mock() as mock:
    mock.foo().returns(10)

print(mock.foo())