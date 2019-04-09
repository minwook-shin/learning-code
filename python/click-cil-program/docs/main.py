import click


@click.command()
def hello():
    """This script prints hello."""
    click.echo('Hello World!')


if __name__ == '__main__':
    hello()
