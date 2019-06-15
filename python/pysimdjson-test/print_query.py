import simdjson

with open('sample.json', 'rb') as fin:
    pj = simdjson.ParsedJson(fin.read())

    print(type(pj))
    print(pj.items('.type'))
    print(pj.items('.created_at'))
    print(pj.items('.id'))
    print(pj.items('.actor[]'))
    print(pj.items('.actor.login'))
    print(pj.items('.repo.name'))
    print(pj.items('.public'))
    print(pj.items('.repo.name'))
    print(pj.items('.payload[]'))
    print(pj.items('.payload.commits[].author'))
