import json

import xmltodict

xml = """<?xml version="1.0"?>
  <site>
     <blog number="1">
        <article id="https://minwook-shin.github.io/">
           <author>minwook-shin</author>
           <url>http://minwook-shin.github.io/python-parsing-serializing-html-documents-using-html5lib/</url>
           <publish_date>2019-05-24T00:00:00+00:00</publish_date>
        </article>
        <article id="https://minwook-shin.github.io/">
           <author>minwook-shin</author>
           <url>http://minwook-shin.github.io/python-parse-and-build-css-using-cssutils/</url>
           <publish_date>2019-05-23T00:00:00+00:00</publish_date>
        </article>
        <article id="https://minwook-shin.github.io/">
           <author>minwook-shin</author>
           <url>http://minwook-shin.github.io/python-html-sanitization-text-linkification-using-bleach/</url>
           <publish_date>2019-05-22T00:00:00+00:00</publish_date>
        </article>
        <article id="https://minwook-shin.github.io/">
           <author>minwook-shin</author>
           <url>http://minwook-shin.github.io/python-iterating-searching-modifying-html-using-beautifulsoup/</url>
           <publish_date>2019-05-21T00:00:00+00:00</publish_date>
        </article>
     </blog>
  </site>"""

import pprint

pp = pprint.PrettyPrinter(indent=4).pprint


result = json.dumps(xmltodict.parse(xml))
pp(result)


def handle_function(_, blog):
    print(blog['url'])
    return True


result = json.dumps(xmltodict.parse(xml, item_depth=3, item_callback=handle_function))
pp(result)

print(type(xmltodict.parse(xml)))
