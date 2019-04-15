from configobj import ConfigObj


config = ConfigObj("config.ini")


name = config['name']
print(name)
nick = config['nickname']
print(nick)


user = config['USER']
user_id = user['id']
user_email = user['email']
print(user)


title = config['blog']['title']
print(title)
body = config['blog']['body']
print(body)

print(config['blog']['tag'])

print(config['visitor'])