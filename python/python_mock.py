from unittest.mock import MagicMock, Mock

class testClass():
    def method(parameter_list):
        pass

test_class = testClass()
test_class.method = MagicMock(return_value=10)
test = test_class.method(1, 2, 3, key='value')
print(test)


mock = Mock(side_effect=KeyError('key Error'))
try:
    mock()
except KeyError as e:
    print(e)


from unittest.mock import patch

with patch.object(testClass, 'method', return_value=None) as mock_method:
    thing = testClass()
    thing.method(1, 2, 3)
print(mock_method)

test_dict = {'key': 'value'}
original = test_dict.copy()
with patch.dict(test_dict, {'newkey': 'new_value'}, clear=True):
    print(test_dict)
print(test_dict)


mock = MagicMock()
mock.__str__.return_value = 'str_magic_value'
print(str(mock))


mock = Mock()
mock.__str__ = Mock(return_value='str_mock_value')
print(str(mock))

from unittest.mock import create_autospec
def function(a):
    pass

mock_function = create_autospec(function, return_value='autospec')
print(mock_function(1))
