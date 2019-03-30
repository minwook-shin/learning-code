import socket


sock = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sock.bind(('127.0.0.1', 5005))
sock.listen()
conn, addr = sock.accept()
while True:
    print(addr)
    data = conn.recv(1024).decode()
    if not data:
        break
    capitalizedSentence = data.upper()
    conn.send(capitalizedSentence.encode())

sock.close()
