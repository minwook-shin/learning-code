import mistune

markdown = mistune.markdown('python **hello, world**!')
print(markdown)


renderer = mistune.Renderer(escape=True, hard_wrap=True)
markdown = mistune.Markdown(renderer=renderer)
markdown_text = markdown("python **hello, world**!")

print(markdown_text)