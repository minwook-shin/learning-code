from flanker.addresslib import address

ap = address.parse('Example example@example.com')
print(ap)

not_email = address.parse('Example @example.com')
print(not_email)

multi_address = address.parse_list('example1@example.com, example2@example.com')
print(multi_address)

multi_address2 = address.parse_list('example1@example.com, example2@example.com', as_tuple=True)
print(multi_address2)

multi_address3 = address.parse_list('example1@example.com, example2@example.com', strict=True)
print(multi_address3)

from flanker.addresslib import validate

sa = validate.suggest_alternate('example@gmail..com')
print(sa)


msg = '''MIME-Version: 1.0
Content-Type: text/plain
From: Example1 <example1@example.com>
To: Example2 <example2@example.com>
Subject: hello, message
Date: Mon, 10 Sep 2019 12:43:03 -0700

this is a single part message.'''

from flanker import mime

fs = mime.from_string(msg)
print(fs.body)
print(fs.headers.items())

print(fs.content_type)
print(fs.subject)

print(fs.content_type.is_singlepart())
print(fs.content_type.is_multipart())


from flanker.mime import create

message = create.text("plain", "hello, world!")
message.headers['From'] = u'Example1 <example1@example.com>'
message.headers['To'] = u'Example2 <example2@example.com>'
message.headers['Subject'] = u"hello"
message = create.from_string(message.to_string())
print(message.to_string())