import webview


if __name__ == '__main__':
    webview.create_window('Debug window',
                          'https://www.google.com',
                          debug=True)