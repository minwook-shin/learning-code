from pyquery import PyQuery

d = PyQuery('<p class="hello" id="hello">hello, world!</p>')
d('p').append(' check out <a href="https://www.google.com">google</a>')
print(d)

p = d('p')
d = PyQuery('<html><body><div id="test"><a href="https://github.com">github</a></div></body></html>')
p.prependTo(d('#test'))
print(d)
p.insertAfter(d('#test'))
print(d)
p.insertBefore(d('#test'))
print(d)

print(PyQuery('<div>hello,world!</div>').addClass('test') + PyQuery('<b>b</b>'))