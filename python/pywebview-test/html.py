import webview
import threading


def load_html():
    webview.load_html('<h1>loaded HTML</h1>')


if __name__ == '__main__':
    t = threading.Thread(target=load_html)
    t.start()

    webview.create_window('Load HTML Example')