from pyquery import PyQuery


d = PyQuery("<option value='1'><option value='2'>")
print(d('option[value="1"]'))


p = PyQuery('<p id="hello" class="hello"></p>')
p.attr.id = "hello_world"
print(p)

p.attr["id"] = "hello_world"
print(p)

p.attr(id='hello', class_='hello2')
print(p)
