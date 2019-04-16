import configparser


config = configparser.ConfigParser()

config.read('config.ini')
sec = config.sections()
print(sec)

example = config['example.com']
ex1 = example['Id']
print(ex1)
ex2 = example['Email']
print(ex2)

user = config['example.org']['User']
print(user)
email = config['example.com']['Email']
print(email)

title = config['DEFAULT']['Title']
print(title)

print('example.org' in config)
print('example.com' in config)

for i in config['example.com']:
    print("for : ", i, " | ", config['example.com'][i])
