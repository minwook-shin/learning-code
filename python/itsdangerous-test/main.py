from itsdangerous import URLSafeSerializer


auth_s = URLSafeSerializer("secret key", "auth")
token = auth_s.dumps({"id": 1, "name": "name"})
print(token)

data = auth_s.loads(token)
print(data["name"])
