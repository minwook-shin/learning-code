class Test(object):
    def __init__(self):
        self.hello: str = ""

    def __get__(self, instance, e) -> str:
        return instance.hello

    def __set__(self, instance, value: str) -> None:
        instance.hello: str = value

    def testStrFunction(self) -> str:
        return "hello, world!"

    def testIntFunction(self) -> int:
        return 10


t = Test()

t.hello = 'hello,world!'
print(t.hello)
