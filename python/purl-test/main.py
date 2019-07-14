from purl import URL

str_url = URL('https://www.google.com/search?q=google')
print(str_url)
print(str_url.as_string())
argument_url = URL(scheme='https', host='www.google.com', path='/search', query='q=google')
print(argument_url)
print(argument_url.as_string())
inline_url = URL().scheme('https').domain('www.google.com').path('search').query_param('q', 'google')
print(inline_url)
print(inline_url.as_string())

u = URL('postgres://username:password@localhost:1234/test?ssl=true')
print(u.scheme())
print(u.host())
print(u.domain())
print(u.username())
print(u.password())
print(u.netloc())
print(u.port())
print(u.path())
print(u.query())
print(u.path_segments())
print(u.query_param('ssl'))
print(u.query_param('ssl', as_list=True))
print(u.query_params())
print(u.has_query_param('ssl'))
print(u.subdomains())

u = URL.from_string('https://github.com/minwook-shin')
print(u.path_segment(0))


new_url = u.add_path_segment('minwook-shin.github.com')
print(new_url.as_string())

from purl import expand
print(expand(u"{/path*}", {'path': ['sub', 'index']}))


from purl import Template
template = Template("http://example.com{/path*}")
url = template.expand({'path': ['sub', 'index']})
print(url.as_string())

