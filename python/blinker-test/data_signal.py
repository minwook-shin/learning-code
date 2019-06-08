from blinker import signal


send_data = signal('send-data')
@send_data.connect
def receive_data(sender, **kw):
    print(sender, kw)


result = send_data.send('anonymous', value="value")