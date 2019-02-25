from typing import *

text: str = "hello, world!"
number: int = 10


def func1(url: str) -> str:
    return "hello, world!"


def func2(url: str) -> None:
    print("hello, world!")


Vector = List[float]

dictionary = Dict[str, int]
tupleArr = Tuple[str, int]

UserId = NewType('UserId', int)
some_id = UserId(524313)


class custom(NamedTuple):
    id: int
    name: str


user = custom(0, "user")


def callable(next_item: Callable[[], str]) -> None:
    pass


T = TypeVar('T')


def GenericFunc(l: Sequence[T]) -> T:
    return l[0]


a = None
b = Any


def foo(bar: Any) -> int:
    return 1


data = ""


def legacy(text: Any) -> Any:
    return data


def unionFunc(p: Union[int, float]) -> float:
    return p


def vec2(x: T, y: T) -> List[T]:
    return [x, y]
