import lassie
from pprint import pprint

sample = lassie.fetch('https://www.youtube.com/watch?v=R6IT_f0XPT8')
print(sample)
pprint(sample)

print("*" * 100)


sample = lassie.fetch('https://www.youtube.com/watch?v=R6IT_f0XPT8', all_images=True)
print(sample)
pprint(sample)

print("*" * 100)


from lassie import Lassie

l = Lassie()
sample = l.fetch('https://www.youtube.com/watch?v=R6IT_f0XPT8')

print(sample)
pprint(sample)

print("*" * 100)

l.request_opts = {
    'headers': {
        'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/605.1.15 (KHTML, like Gecko) '
                      'Version/12.1.1 Safari/605.1.15 '
    }
}

l.request_opts = {
    'timeout': 0.1
}

sample = l.fetch('https://www.youtube.com/watch?v=R6IT_f0XPT8')

print(sample)
pprint(sample)