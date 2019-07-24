from htmlparsing import Attr, Text
from toapi import Api, Item

api = Api()


@api.site('https://minwook-shin.github.io')
@api.list('.post')
@api.route('/api', '/')
class Post(Item):
    url = Attr('.read-more', 'href')
    title = Text('h1 > a')


api.run(debug=True, host='0.0.0.0', port=5000)