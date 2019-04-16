import configparser


config = configparser.ConfigParser()

config['example.org'] = {}
config['example.org']['User'] = 'user_name'

config['example.com'] = {}
example = config['example.com']
example['ID'] = 'real_id'
example['Email'] = 'example@example.com'

config['DEFAULT']['Agree'] = 'True'

config['DEFAULT'] = {'Title': 'hello, world!',
                     'Body': 'print(hello,world!)', 'Author': 'realName'}
with open('config.ini', 'w') as configfile:
    config.write(configfile)
