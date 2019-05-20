from bs4 import BeautifulSoup

html_content = """
<html><head><title>html title</title></head>
<body>
<p class="title"><b>hello, world!</b></p>

<p class="story">this is first story</p>
<a href="http://example.com" id="link1">test</a>,
<a href="http://example.com" id="link2">test</a>,
<p class="story">this is second story</p>
"""

soup = BeautifulSoup(html_content, 'html.parser')


print(soup.prettify())


print("title:", soup.title)

print("title name:", soup.title.name)

print("title string:", soup.title.string)

print("title parent:", soup.title.parent.name)

print("p:", soup.p)

print("p class:", soup.p['class'])

print("a:", soup.a)

for url in soup.find_all('a'):
    print(url)
for url in soup.find_all('a'):
    print(url.get("href"))
print("all a:", soup.find_all('a'))

print("a link1:", soup.find(id="link1"))

print(soup.get_text())