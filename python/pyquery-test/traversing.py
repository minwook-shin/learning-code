from pyquery import PyQuery

d = PyQuery('<p id="hello" class="hello"><a/></p><p id="test"><a/></p>')
print(d('p').filter('.hello'))

print(d('p').find('a'))
print(d('p').eq(1).find('a'))

print(d('p').find('a').end())

d = PyQuery('<p id="hello.you"><a/></p><p id="test"><a/></p>')
print(d(r'#hello\.you'))
