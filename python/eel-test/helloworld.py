import eel

eel.init('web', allowed_extensions=['.js', '.html'])


@eel.expose
def say_hello_py(x):
    print('Hello from %s' % x)


say_hello_py('Python World!')
eel.say_hello_js('Python World!')

eel.start('helloworld.html', size=(300, 200))