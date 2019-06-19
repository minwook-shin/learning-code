import pyexcel

sheet = pyexcel.Sheet()

sheet.name = "name"
sheet.ndjson = """
{"year": ["2017", "2018", "2019", "2020", "2021"]}
{"user": [129, 253, 304, 403, 545]}
{"visit": [203, 403, 632, 832, 1023]}
""".strip()

print(sheet)
print(type(sheet))
