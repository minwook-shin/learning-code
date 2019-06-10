import subprocess


lst = subprocess.run(["ls", "-l"])
print(lst.stdout)


import delegator


lst2 = delegator.run('ls', "-l")
print(lst2.out)


c = delegator.run(['curl', 'www.google.com'], block=False)

print(c.pid)
c.block()
print(c.out)
print(c.return_code)

try:
    c = delegator.chain('fortune | cowsay')
    print(c.out)
except FileNotFoundError:
    print("맥 : brew install cowsay / brew install fortune")

c = delegator.chain('env | grep NEWENV', env={'NEWENV': 'FOO_BAR'})
print(c.out)
