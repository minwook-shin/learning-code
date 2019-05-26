from pyquery import PyQuery

d = PyQuery('<div><input type="button"/><button></button></div>')
print(d(':button'))

d = PyQuery('<div><input type="file"/><textarea></textarea></div>')
print(d(':input'))

d = PyQuery('<div><h1/><h1 class="title">title</h1></div>')
print(d('h1:contains("title")'))

d = PyQuery('<div><h1><span>title</span></h1><h1/></div>')
print(d('h1:parent'))

d = PyQuery('<div><p class="first"></p><p></p></div>')
print(d('p:first'))
print(d('p:last'))

d = PyQuery('<div><h1><span>title</span></h1><h2/></div>')
print(d(':empty'))

d = PyQuery('<div><h1>title</h1></div>')
print(d(':header'))

d = PyQuery('<div class="foo"><div class="bar"></div></div>')
print(d('.foo:has(div)'))
