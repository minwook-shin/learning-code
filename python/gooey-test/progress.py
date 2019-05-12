import sys
from time import sleep
from gooey import Gooey, GooeyParser


@Gooey(progress_regex=r"^\[progress\] (\d+)$")
# disable_stop_button=True
def main():
    parser = GooeyParser(prog="progress_bar example")
    _ = parser.parse_args(sys.argv[1:])

    for i in range(100):
        print("[progress] {}".format(i + 1))
        sys.stdout.flush()
        sleep(0.1)


if __name__ == "__main__":
    main()