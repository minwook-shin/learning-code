import webview
import threading


def toggle_fullscreen():
    webview.toggle_fullscreen()


if __name__ == '__main__':
    t = threading.Thread(target=toggle_fullscreen)
    t.start()

    webview.create_window("Full-screen window",
                          "https://www.google.com")