import bleach

tag = bleach.clean('an <script>script()</script> example')
print(tag)

url = bleach.linkify('an http://example.com url')
print(url)