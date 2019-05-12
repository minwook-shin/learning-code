import eel

eel.init('web')


@eel.expose
def handle_input(x):
    print('%s' % x)


eel.say_hello_js('connected!')

eel.start('input.html', size=(500, 200))