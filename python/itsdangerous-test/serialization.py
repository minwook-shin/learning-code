from itsdangerous.serializer import Serializer
s = Serializer("secret-key")
serializer_value = s.dumps([1, 2, 3, 4])
print(serializer_value)
print(s.loads(serializer_value))


from itsdangerous.url_safe import URLSafeSerializer
s1 = URLSafeSerializer("secret-key", salt="activate")
s1.dumps(1)
s2 = URLSafeSerializer("secret-key", salt="upgrade")
s2.dumps(1)
print(s1.dumps(1))
print(s2.dumps(1))
print(s1.loads(s1.dumps(1)))
print(s2.loads(s2.dumps(1)))


s = URLSafeSerializer("secret-key")
s.dumps([1, 2, 3, 4])
print(s.dumps([1, 2, 3, 4]))
print(s.loads(s.dumps([1, 2, 3, 4])))