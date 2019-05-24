import html5lib


document1 = html5lib.parse("<p>Hello World!</p>")
print(document1)


from urllib.request import urlopen

with urlopen("http://www.google.com/") as f:
    document2 = html5lib.parse(f, transport_encoding=f.info().get_content_charset())
    print(document2)


document3 = html5lib.HTMLParser(tree=html5lib.getTreeBuilder("dom")).parse("<p>Hello World!</p>")
print(document3)

element = html5lib.parse('<p>Hello World!</p>')
walker = html5lib.getTreeWalker("etree")
stream = walker(element)
s = html5lib.serializer.HTMLSerializer().serialize(stream)
for i in s:
    print(i)


from html5lib.filters import sanitizer

dom = html5lib.parse("<script>alert('warning!')</script>", treebuilder="dom")
walker = html5lib.getTreeWalker("dom")
clean_stream = sanitizer.Filter(walker(dom))
print(clean_stream)