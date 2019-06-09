from itsdangerous import Signer, BadSignature

s = Signer("secret-key")
signed_text = s.sign("test string")
print(s.unsign(signed_text))

try:
    s.unsign(b"different string.wh6tMHxLgJqB6oY1uT73iMlyrOA")
except BadSignature:
    print("Signature does not match")


from itsdangerous import TimestampSigner, SignatureExpired
import time

s = TimestampSigner('secret-key')
signed_text = s.sign('text')
time.sleep(2)
try:
    print(s.unsign(signed_text, max_age=1))
except SignatureExpired:
    print("SignatureExpired")