import uuid


host_uuid = uuid.uuid1()
print(host_uuid)

random_uuid = uuid.uuid4()
print(random_uuid)

url = "http://www.google.com"

md5_url_uuid = uuid.uuid3(uuid.NAMESPACE_URL, url)
print(md5_url_uuid)

sha_1_url_uuid = uuid.uuid5(uuid.NAMESPACE_URL, url)
print(sha_1_url_uuid)
