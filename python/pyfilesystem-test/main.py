from fs.osfs import OSFS


home_fs = OSFS("~/")
print(home_fs)

dir_list = home_fs.listdir('/')
print(dir_list)

from fs import open_fs


home_fs = open_fs('osfs://~/')
print(home_fs.listdir('/'))

current_fs = open_fs('.')
print(current_fs.listdir('/'))

my_fs = open_fs('.')
print(my_fs.tree())


current_fs = open_fs('.')
current_fs.writetext('reminder.txt', 'buy pycharm')
current_fs.close()

with open_fs('.') as home_fs:
    home_fs.writetext('reminder.txt', 'buy coffee')


with open_fs('.') as home_fs:
    with home_fs.open('reminder.txt') as reminder_file:
        print(reminder_file.read())


current_fs = open_fs('.')
print(current_fs.listdir('/.idea'))

current_fs = open_fs('.')
directory = list(current_fs.scandir('/.idea'))
for i in directory:
    print(i)

current_fs = open_fs('.')
if not current_fs.exists('text'):
    folder_fs = home_fs.makedirs('text')
    folder_fs.touch('__init__.py')
    folder_fs.writetext('README.md', "## hello,world!")
    print(folder_fs.listdir('/'))

current_fs = open_fs('.')
for path in current_fs.walk.files(filter=['*.py']):
    print(path)


from fs.copy import copy_fs
copy_fs('./text', 'zip://text.zip')