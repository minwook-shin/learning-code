from paramiko.py3compat import input

import paramiko
import getpass
import sys


paramiko.util.log_to_file("logging.log")

hostname = input("Host: ")

if len(hostname) == 0:
    sys.exit(1)

if hostname.find(":") >= 0:
    hostname, port_str = hostname.split(":")
    port = int(port_str)

default_username = getpass.getuser()

username = input("Username [%s]: " % default_username)
if len(username) == 0:
    username = default_username

password = getpass.getpass("Password: ")

client = paramiko.SSHClient()

client.load_system_host_keys()

client.set_missing_host_key_policy(paramiko.WarningPolicy())

client.connect(hostname, 22, username, password)

client.close()
