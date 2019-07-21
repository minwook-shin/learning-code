from micawber.providers import Provider, bootstrap_oembed
from pprint import pprint

youtube = Provider('http://www.youtube.com/oembed')
test = youtube.request('https://www.youtube.com/watch?v=R6IT_f0XPT8')
pprint(test)


"""
    {
        "provider_name": "YouTube",
        "provider_url": "https:\/\/www.youtube.com\/",
        "endpoints": [
            {
                "schemes": [
                    "https:\/\/*.youtube.com\/watch*",
                    "https:\/\/*.youtube.com\/v\/*",
                    "https:\/\/youtu.be\/*"
                ],
                "url": "https:\/\/www.youtube.com\/oembed",
                "discovery": true
            }
        ]
    },
"""
providers = bootstrap_oembed()
test = providers.request('https://www.youtube.com/watch?v=R6IT_f0XPT8')
pprint(test)