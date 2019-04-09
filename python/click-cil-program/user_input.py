import click

value = click.prompt('Please enter a valid integer', type=int)

try:
    click.confirm('Do you want to continue?', abort=True)
except click.exceptions.Abort as e:
    pass
