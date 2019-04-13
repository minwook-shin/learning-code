import eventlet


def handle(fd):
    print("connected")
    while True:
        x = fd.readline()
        if not x:
            break
        fd.write(x)
        fd.flush()
        print(x, end=' ')
    print("disconnected")


server = eventlet.listen(('127.0.0.1', 5000))
pool = eventlet.GreenPool()
while True:
    new_sock, address = server.accept()
    print("accepted", address)
    pool.spawn_n(handle, new_sock.makefile('rw'))
