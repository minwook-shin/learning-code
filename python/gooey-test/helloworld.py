from gooey import Gooey, GooeyParser


@Gooey()
def main():
    parser = GooeyParser(description='HELLO WORLD.')

    parser.add_argument(
        'string_field',
        metavar='enter string',
        help='hello, world!')

    args = parser.parse_args()
    print(args.string_field)


if __name__ == '__main__':
    main()
