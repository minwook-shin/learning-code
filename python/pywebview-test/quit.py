import webview


if __name__ == '__main__':
    webview.create_window("Confirm Quit Example",
                          "https://www.google.com",
                          confirm_quit=True)