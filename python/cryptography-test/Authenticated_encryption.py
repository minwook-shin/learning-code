import os
from cryptography.hazmat.primitives.ciphers.aead import AESGCM


data = b"a secret message"

key = AESGCM.generate_key(bit_length=128)
os_urandom = os.urandom(12)

encrypt_string = AESGCM(key).encrypt(os_urandom, data, b"authenticated")
print(encrypt_string)

decrypt_string = AESGCM(key).decrypt(os_urandom, encrypt_string, b"authenticated")
print(decrypt_string)