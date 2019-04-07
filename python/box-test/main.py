from box import Box

user_data = {
    "users": {
        "Kim": {
            "id": "realKimID",
            "email": "Kim@example.com",
            "friend": [{"name": "Lee", "status": "activated"},
                       {"name": "ahn", "status": "activated"}]
        },
        "Shin": {
            "id": "realShinID",
            "email": "Shin@example.com",
            "friend": [{"name": "Lee", "status": "deactivated"},
                       {"name": "ahn", "status": "deactivated"}]
        },
    }
}

user_box = Box(user_data)

print(user_box.users.Kim.id)

print(user_box.users.Kim.friend[0])

user_box.users.Kim.friend.append({"name": "han", "status": "activated"})

print(user_box.users.Kim.friend[-1])

print(user_box.users.Kim.to_dict())
print(user_box.users.Kim == user_box.users.Kim.to_dict())

test = Box({'hello': 10}, world=10)
print(test)

test2 = Box({'hello': 10}, world=10)
print(test2.hello == test2['hello'] == getattr(test2, 'hello'))


frozen_box = Box(data={'frozen': 'box',}, frozen_box=True)
# frozen_box.frozen = 'Box'
# box.BoxError: Box is frozen

empty_box = Box(default_box=True)

print(empty_box.a.b.c)
# <Box: {}>

empty_box.a.b.c.d = "h"
print(empty_box)

camel_killer_box = Box(HelloWorld="HELLO, WORLD!", camel_killer_box=True)
print(camel_killer_box.hello_world)

box_of_order = Box(ordered_box=True)
box_of_order.c = 1
box_of_order.a = 2
box_of_order.d = 3
print(box_of_order.keys() == ['c', 'a', 'd'])


from box import BoxList
box_list = BoxList({'item': x} for x in range(10))

print(box_list[5].item)


from box import SBox
s_box = SBox(hello='world')
print(s_box.json)
