from prompt_toolkit.shortcuts import message_dialog, input_dialog, yes_no_dialog, button_dialog

message_dialog(
    title='Example dialog window',
    text='Do you want to continue?\nPress ENTER to quit.')

result = yes_no_dialog(
    title='Example dialog window',
    text='Do you want to continue?')

input_dialog(
    title='Example dialog window',
    text='input:')

button_dialog(
    title='Example dialog window',
    text='Do you want to continue??',
    buttons=[
        ('Yes', True),
        ('No', False),
    ],
)