from prompt_toolkit import print_formatted_text

print_formatted_text('Hello world')


from prompt_toolkit import print_formatted_text, HTML

print_formatted_text(HTML('<b>bold</b>'))
print_formatted_text(HTML('<i>talic</i>'))
print_formatted_text(HTML('<u>underline</u>'))