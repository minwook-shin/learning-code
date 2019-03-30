import socket

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
while True:
    msg = input("->")
    sock.sendto(msg.encode(), ("127.0.0.1", 5005))
    recvMsg, addr = sock.recvfrom(2048)
    print(recvMsg.decode())

sock.close()
