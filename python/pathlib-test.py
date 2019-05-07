from pathlib import PureWindowsPath, WindowsPath
from pathlib import PurePath, PurePosixPath
from pathlib import Path


p = Path('.')
for x in p.iterdir():
    print(x)

print("*"*100)

p = Path('.')
for x in p.iterdir():
    if x.is_dir():
        print(x)

print("*"*100)

p = Path('.')
for x in p.iterdir():
    if not x.is_dir():
        print(x)

p = Path('.')
pyList = list(p.glob('*.py'))
print(pyList)


p = Path('/etc')
print(p)

p = Path('/etc')
exist = p.exists()
isDir = p.is_dir()
print(exist)
print(isDir)

p = Path('./python-rx-basic.py')
with p.open() as f:
    print(f.readline())


cwd = Path.cwd()
print(cwd)

home = Path.home()
print(home)

pure = PurePath('setup.py')
print(pure)

ex1 = PurePath('a', 'b/c', 'd')
print(ex1)
ex2 = PurePosixPath('a/b/c/d')
print(ex2)
ex3 = PurePath(Path('a'), Path('b'), Path('c'), Path('d'))
print(ex3)
print(ex1 == ex2 == ex3)

print(PurePosixPath('/etc'))
print(PureWindowsPath('c:/Windows'))
try:
    print(WindowsPath('setup.py'))
except NotImplementedError as e:
    print("cannot instantiate 'WindowsPath' on your system")
