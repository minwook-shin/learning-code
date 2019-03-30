import socket

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)  # UDP
sock.bind(("127.0.0.1", 5005))
while True:
    data, addr = sock.recvfrom(1024)
    data = data.decode().upper()
    print(data)
    sock.sendto(data.encode(), addr)
sock.close()
