import click


@click.group()
def cli():
    pass


@click.command()
def add_user():
    click.echo('Added User.')


cli.add_command(add_user)

if __name__ == '__main__':
    cli()
