import bleach

tag1 = bleach.clean('<b><i>example</i></b>', tags=['b'])
print(tag1)

tag2 = bleach.clean('<p class="foo" style="color: red; font-weight: bold;">example</p>',
                    tags=['p'],
                    attributes=['style'],
                    styles=['color'])
print(tag2)

attrs = {'a': ['href', 'rel'],'img': ['alt']}

tag3 = bleach.clean('<img alt="an example" width=500>', tags=['img'], attributes=attrs)
print(tag3)


def allow_h(_, name, __):
    return name[0] == 'h'


tag4 = bleach.clean('<a href="http://example.com" title="link">example</a>', tags=['a'], attributes=allow_h)
print(tag4)

tag5 = bleach.clean('<a href="smb://more_text">example</a>',protocols=['http', 'https', 'smb'])
print(tag5)

tag6 = bleach.clean('<span>example</span>', strip=True)
print(tag6)


html = '<!-- commented -->example'

tag7 = bleach.clean(html)
print(tag7)