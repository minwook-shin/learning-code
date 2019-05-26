import xmldataset

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


profile = """
site
    blog
        number     = external_dataset:blog_information
        article
            id     = dataset:blog_article,prefix:blog_article_
            author = dataset:blog_article,prefix:blog_article_
            url    = dataset:blog_article,prefix:blog_article_
            publish_date  = dataset:blog_article,name:date,prefix:blog_article_
            __EXTERNAL_VALUE__ = blog_information:number:blog_article"""

result = xmldataset.parse_using_profile(xml, profile)

import pprint

pp = pprint.PrettyPrinter(indent=4).pprint

pp(result)

print(type(result))