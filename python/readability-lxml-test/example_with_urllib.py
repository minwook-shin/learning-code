import urllib.request
from readability import Document

req = urllib.request.Request(
    'https://en.wikipedia.org/wiki/%22Hello,_World!%22_program',
    headers={
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.1 Safari/605.1.15'
    }
)

with urllib.request.urlopen(req) as f:
    urllib_content = f.read()
# print(urllib_content.decode("utf-8"))

doc = Document(urllib_content)

print(doc.title())
print(doc.short_title())

print(doc.summary())