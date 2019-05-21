import bleach
from bleach.linkifier import Linker

link1 = bleach.linkify('http://example.com example')
print(link1)


def set_title(attrs, _):
    attrs[(None, 'title')] = 'example title'
    return attrs


linker = Linker(callbacks=[set_title])
link2 = linker.linkify('http://example.com example')
print(link2)


def allowed_attrs(attrs, _):
    allowed = [
        (None, 'href'),
        (None, 'style'),
        '_text',
    ]
    return dict((k, v) for k, v in attrs.items() if k in allowed)


linker = Linker(callbacks=[allowed_attrs])
link3 = linker.linkify('<a style="font-weight: super bold;" href="http://example.com">example</a>')
print(link3)


def shorten_url(attrs, _):
    text = attrs['_text']
    if len(text) > 20:
        attrs['_text'] = text[0:15] + '...'
        return attrs


linker = Linker(callbacks=[shorten_url])
link4 = linker.linkify('http://example.com/longlongurl')
print(link4)


def remove_http(attrs, _):
    if attrs[(None, 'href')].startswith('http:'):
        return None


linker = Linker(callbacks=[remove_http])
link5 = linker.linkify('<a href="http://example.com">example</a>')
print(link5)