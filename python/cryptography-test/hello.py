from cryptography.fernet import Fernet


key = Fernet.generate_key()

token = Fernet(key).encrypt(b"hello, world")
print(token)

decrypt_string = Fernet(key).decrypt(token)
print(decrypt_string)
