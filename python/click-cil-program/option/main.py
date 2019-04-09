import click


@click.command()
@click.option('--url', '-u', 'arg', required=True, type=str)
def url_check(arg):
    click.echo('URL: %s' % arg)


if __name__ == '__main__':
    url_check()
