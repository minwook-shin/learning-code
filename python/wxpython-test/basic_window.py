import wx


def print_hello(event):
    wx.MessageBox("Hello world!")


def about_func(event):
    wx.MessageBox("sample text", "About title", wx.OK | wx.ICON_INFORMATION)


class HelloFrame(wx.Frame):
    def __init__(self, *args, **kw):
        super(HelloFrame, self).__init__(*args, **kw)

        pnl = wx.Panel(self)

        wx.StaticText(pnl, label="Hello World!")

        file = wx.Menu()
        hello_item = file.Append(-1, "&Hello...\tCtrl-H", "hello!")

        help_menu = wx.Menu()
        about_item = help_menu.Append(wx.ID_ABOUT)

        menu_bar = wx.MenuBar()
        menu_bar.Append(file, "&File")
        menu_bar.Append(help_menu, "&Help")

        self.SetMenuBar(menu_bar)

        self.Bind(wx.EVT_MENU, print_hello, hello_item)
        self.Bind(wx.EVT_MENU, about_func, about_item)

        self.CreateStatusBar()
        self.SetStatusText("wxPython status text")


if __name__ == '__main__':
    app = wx.App()
    frm = HelloFrame(None, title='basic window')
    frm.Show()
    app.MainLoop()