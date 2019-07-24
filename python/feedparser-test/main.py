import feedparser
from pprint import pprint

data = """<rss version="2.0">
<channel>
<title>test feed</title>
</channel>
</rss>"""
local_fees = feedparser.parse(data)
pprint(local_fees["feed"])
pprint(local_fees.feed)
pprint(local_fees.feed.title)

feed = feedparser.parse('feed:https://minwook-shin.github.io/feed')
pprint(feed["feed"])
pprint(feed.feed)

pprint(feed["entries"])
pprint(feed.entries[0].title)
pprint(feed.entries[0].link)
pprint(feed.entries[0].description)
pprint(feed.entries[0].published)
pprint(feed.entries[0].published_parsed)
pprint(feed.entries[0].id)