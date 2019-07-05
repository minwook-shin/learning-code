import difflib
import sys

original = ['python\n', 'java\n', 'go\n']
edited = ['pythonic\n', 'javachip\n', 'go\n']
sys.stdout.writelines(difflib.context_diff(original, edited, fromfile='original.py', tofile='edit.py'))

n = difflib.ndiff("""
python
java
go""".splitlines(keepends=True),

"""
pythonic
javachip
go""".splitlines(keepends=True))
print(''.join(n), end="")

import difflib

n = difflib.ndiff("""
python
java
go""".splitlines(keepends=True),

"""
pythonic
javachip
go""".splitlines(keepends=True))
diff = list(n)
print(''.join(difflib.restore(diff, 1)), end="")
print(''.join(difflib.restore(diff, 2)), end="")

import difflib
import sys

original = ['python\n', 'java\n', 'go\n']
edited = ['pythonic\n', 'javachip\n', 'go\n']
sys.stdout.writelines(difflib.unified_diff(original, edited, fromfile='before.py', tofile='after.py'))

import difflib

text_1 = '''
#include <iostream>
int main() 
{
    return 0;
}'''.splitlines(keepends=True)

text_2 = '''
#include <iostream>
using namespace std;
int main() 
{
    cout << "Hello, World!";
    return 0;
}'''.splitlines(keepends=True)

diff = difflib.Differ()

result = list(diff.compare(text_1, text_2))
print(result)