from dataclasses import dataclass, field, asdict, astuple, make_dataclass
from typing import List


@dataclass
class User:
    name: str = ""
    agree: bool = True
    tag: List[str] = field(default_factory=list)


p = User("Kim", True)
print(asdict(p))
print(type(p))
print(type(asdict(p)))
assert asdict(p) == {'name': "Kim", 'agree': True,'tag': []}

print(astuple(p))
print(type(p))
print(type(astuple(p)))
assert astuple(p) == ("Kim", True, [])

Article = make_dataclass('article',
                   [('title', str),
                    ('text', str),
                    ('index', int, field(default=5))])

a = Article("title","text",1)
print(a)
print(asdict(a))
print(astuple(a))