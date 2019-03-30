import socket

sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.connect(('127.0.0.1', 5005))
while True:
    msg = input("->")
    sock.send(msg.encode())
    data = sock.recv(1024)
    print(data.decode())

sock.close()
