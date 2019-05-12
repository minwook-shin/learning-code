from gooey import Gooey, GooeyParser


@Gooey(program_name="Widget Demo")
def main():
    parser = GooeyParser(description="Gooey's widgets Example ")

    parser.add_argument("directory")
    parser.add_argument(
        "FileChooser", widget="FileChooser")
    parser.add_argument(
        "DirectoryChooser", widget="DirChooser")
    parser.add_argument(
        "FileSaver", widget="FileSaver")
    parser.add_argument(
        "MultiFileSaver", widget="MultiFileChooser")

    parser.add_argument('-c', '--count', action='count')
    parser.add_argument("-s", "--store-true", action="store_true")
    parser.add_argument("-i", "--input", default="default string")
    parser.add_argument('-t', '--type', default=0, type=int)
    parser.add_argument('-d', '--dateChooser', type=int, widget='DateChooser')
    parser.add_argument('-r', '--recursive', choices=['yes', 'no'])

    group = parser.add_mutually_exclusive_group()
    group.add_argument('-g1', '--group-1', dest='group1', action="store_true")
    group.add_argument('-g2', '--group-2', dest='group2', action="store_true")

    args = parser.parse_args()


if __name__ == '__main__':
    main()