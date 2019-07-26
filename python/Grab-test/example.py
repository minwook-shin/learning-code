from grab import Grab


g = Grab()
resp = g.go('http://example.com/')

print(resp.unicode_body())

print(resp.charset)

print(resp.body)

g = Grab()
g.go('http://testing-ground.scraping.pro/login')
g.doc.set_input('usr', 'admin')
g.doc.set_input('pwd', '12345')
g.submit()
print(g.doc.body)


