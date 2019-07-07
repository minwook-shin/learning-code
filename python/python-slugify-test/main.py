from slugify import slugify


txt = 'hello &amp; world!'
r = slugify(txt)
print(r)

txt = 'you are #1'
r = slugify(txt)
print(r)


txt = "___This is a test ---"
r = slugify(txt)
print(r)
txt = "This -- is a ## test ---"
r = slugify(txt)
print(r)

txt = 'one two three four five'
r = slugify(txt, max_length=17)
print(r)
txt = 'one two three four five'
r = slugify(txt, max_length=18, word_boundary=True)
print(r)

txt = 'one two three four five'
r = slugify(txt, max_length=20, word_boundary=True, separator=".")
print(r)


txt = 'one two three four'
r = slugify(txt, max_length=12, word_boundary=True, save_order=True)
print(r)
txt = 'one two three four'
r = slugify(txt, max_length=12, word_boundary=True, save_order=False)
print(r)


txt = 'hello, stop world!'
r = slugify(txt, stopwords=['stop'])
print(r)


txt = "_Passed 1 test_"
regex_pattern = r'[^-a-z_]+'
r = slugify(txt, regex_pattern=regex_pattern)
print(r)

txt = 'python | java & go'
r = slugify(txt, replacements=[['|', 'or'], ['&', 'and']])
print(r)
