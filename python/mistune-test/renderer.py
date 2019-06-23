import re
from mistune import Renderer, InlineLexer, Markdown


class LinkRenderer(Renderer):
    def wiki_document_link(self, alt, link):
        return '<a href="%s">%s</a>' % (link, alt)


class LinkInlineLexer(InlineLexer):
    def enable_link(self):
        self.rules.wiki_document_link = re.compile(r'\[\[([\s\S]+?\|[\s\S]+?)\]\](?!\])')
        self.default_rules.insert(0, 'wiki_document_link')

    def output_wiki_document_link(self, m):
        text = m.group(1)
        print(text)
        alt, link = text.split('|')
        return self.renderer.wiki_document_link(alt, link)


renderer = LinkRenderer()
inline = LinkInlineLexer(renderer)
inline.enable_link()
markdown = Markdown(renderer, inline=inline)
html_text = markdown('[[Link Text|Wiki Link]]')
print(html_text)
