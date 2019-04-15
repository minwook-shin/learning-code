from configobj import ConfigObj

config = ConfigObj()
config.filename = "config.ini"

config['name'] = "name"
config['nickname'] = "nick_name"

config['USER'] = {}
config['USER']['id'] = "real_id"
config['USER']['email'] = "example@example.com"

blog = {
    'title': "hello,world",
    'body': "print(hello,world)",
    'tag': {
        'hello': "world"
    }
}
config['blog'] = blog

config['visitor'] = {}
config['visitor']['today'] = ["1", "2", "3", "4", "5", "6", "7"]

config.write()
