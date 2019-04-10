from prompt_toolkit import prompt

text = prompt('input: ')
print('output: %s' % text)


from prompt_toolkit.auto_suggest import AutoSuggestFromHistory
from prompt_toolkit import PromptSession
from prompt_toolkit.history import FileHistory

session = PromptSession(history=FileHistory('my_history'))

text1 = session.prompt('session input1: ',)
text2 = session.prompt('session input2: ', auto_suggest=AutoSuggestFromHistory())


from prompt_toolkit.completion import WordCompleter
from prompt_toolkit import prompt

word_completer = WordCompleter(['hello', 'world', 'python'])
text = prompt('input: ', completer=word_completer, complete_while_typing=True)
print('output: %s' % text)


from prompt_toolkit.validation import Validator
from prompt_toolkit import prompt


def is_number(txt):
    return txt.isdigit()


validator = Validator.from_callable(
    is_number,
    error_message='non-numeric characters',
    move_cursor_to_end=True)

number = int(prompt('input: ', validator=validator))
print('output: %s' % number)

