import cryptography.exceptions
from cryptography.hazmat.primitives.asymmetric.ed25519 import Ed25519PrivateKey


private_key = Ed25519PrivateKey.generate()

signature = private_key.sign(b"example@example.com")

public_key = private_key.public_key()

try:
    verify_result = public_key.verify(signature, b"example@example.com")
except cryptography.exceptions.InvalidSignature:
    print("Wrong key")