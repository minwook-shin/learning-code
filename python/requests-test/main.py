import requests

r = requests.get('https://api.github.com')
print(r)
print(r.content, type(r.content))
print(r.text, type(r.text))
print(r.url)
print(r.status_code)
print(r.json())
print(r.headers)

r = requests.post('https://httpbin.org/post', data={'key': 'value'})
print(r)
print(r.content, type(r.content))
print(r.text, type(r.text))
print(r.url)
print(r.status_code)
print(r.json())
print(r.headers)

r = requests.put('https://httpbin.org/put', data={'key': 'value'})
print(r)
print(r.content, type(r.content))
print(r.text, type(r.text))
print(r.url)
print(r.status_code)
print(r.json())
print(r.headers)

payload = {'key1': 'value1', 'key2': 'value2'}
r = requests.get('https://httpbin.org/get', params=payload)
print(r)
print(r.content, type(r.content))
print(r.text, type(r.text))
print(r.url)
print(r.status_code)
print(r.json())
print(r.headers)

url = 'https://www.google.com/search?q=python'
headers = {'user-agent': "Mozilla/5.0"}
r = requests.get(url, headers=headers)

r = requests.get('https://httpbin.org/cookies', cookies=dict(hello_cookies='hello,world!'))
print(r.text)
