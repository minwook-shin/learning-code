from pyquery import PyQuery


p = PyQuery('<p id="hello" class="hello"></p>')('p')
print(p.addClass("hello_world"))
print(p.toggleClass("python hello_world"))
print(p.removeClass("python"))

p = PyQuery('<p id="hello" class="hello"></p>')
p.css.font_size = "1px"
print(p.attr.style)
p.css['font-size'] = "2px"
print(p.attr.style)
p.css(font_size="3px")
print(p.attr.style)
p.css = {"font-size": "4px"}
print(p.attr.style)
