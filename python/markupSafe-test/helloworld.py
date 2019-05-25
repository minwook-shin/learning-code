from markupsafe import Markup, escape

escape_code1 = escape('<script>alert(hello world!);</script>')
print(escape_code1)
print(escape_code1.unescape())

escape_code2 = Markup.escape('<script>alert(hello world!);</script>')
print(escape_code2)
print(escape_code2.unescape())

markup_code1 = Markup('<script>alert(hello world!);</script>')
print(markup_code1)

template = Markup("Hello <em>%s</em>")
markup_code2 = template % '"World"'
print(markup_code2)


class Image:
    def __init__(self, url):
        self.url = url

    def __html__(self):
        return '<img src="%s">' % self.url


img = Image('logo.png')
img_code = Markup(img)
print(img_code)


class User:
    def __init__(self, id, name):
        self.id = id
        self.name = name

    def __html__(self):
        return '<a href="/user/{}">{}</a>'.format(
            self.id, escape(self.name)
        )


user = User(1, '<script>')
user_code = escape(user)
print(user_code)