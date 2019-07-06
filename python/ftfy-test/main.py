from ftfy import fix_text, explain_unicode


print(fix_text('uÌˆnicode'))

print(fix_text('&lt;3'))

print(fix_text("&macr;\\_(ã\x83\x84)_/&macr;"))

len(fix_text(''))

explain_unicode('ノ( º _ ºノ) 테스트')

from ftfy.fixes import fix_encoding, unescape_html, uncurl_quotes, fix_line_breaks, decode_escapes


print(fix_encoding('â\x81”.'))

print(unescape_html('&lt;hr&gt;'))

print(uncurl_quotes('\u201ctest\u201d'))

print(fix_line_breaks("1. hello\u2028" "2. world"))

factoid = '\\u20a2'
print(decode_escapes(factoid))


from ftfy.formatting import character_width, display_center


print(character_width('A'))
print(character_width('가'))

lines = ['Display center', 'center']
for line in lines:
    print(display_center(line, 20, '▒'))