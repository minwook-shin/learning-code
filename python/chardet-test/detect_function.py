import chardet

import urllib.request

with urllib.request.urlopen('https://google.com') as f:
    response = f.read()

output = chardet.detect(response)
print(output)

with urllib.request.urlopen('https://naver.com') as f:
    response = f.read()

output = chardet.detect(response)
print(output)


with urllib.request.urlopen('http://www.auction.co.kr') as f:
    response = f.read()

output = chardet.detect(response)
print(output)


from chardet.universaldetector import UniversalDetector

detector = UniversalDetector()
with urllib.request.urlopen('http://www.auction.co.kr') as f:
    response = f
    for line in response.readlines():
        detector.feed(line)
        if detector.done:
            break
detector.close()

print(detector.result)