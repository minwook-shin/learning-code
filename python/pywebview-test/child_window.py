import webview
import threading


def create_window():
    webview.create_window('Window #2',
                          "https://www.google.com")


if __name__ == '__main__':
    t = threading.Thread(target=create_window)
    t.start()
    webview.create_window('Window #1',
                          "https://www.google.com")