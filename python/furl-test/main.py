from furl import furl

f = furl('https://www.google.co.kr/search?source=hp&q=google')
f.args['test'] = 'test'
del f.args['source']
print(f.url)

print(furl('https://www.google.co.kr/search?source=hp&q=google').add({'test': 'test'}).url)

print(furl('https://www.google.co.kr/search?source=hp&q=google').set({'q': 'google'}).url)

print(furl('https://www.google.co.kr/search?source=hp&q=google').remove(['source']).url)


f = furl('http://www.google.com/')
f.path = 'index'
f.args['test'] = 'test'
print(f.url)


f = furl('http://www.google.com/')
f.fragment.path.segments = ['first', 'directories']
f.fragment.args = {'test': 'test'}
print(f.url)

f = furl('http://www.google.com/path?q=google')
print(f.copy().remove(path=True).url)
print(f.copy().set(host='host.com').url)
print(f.copy().join('/index.html').url)
print(f.copy().add(fragment_path='add_fragment').url)


print(f.asdict())