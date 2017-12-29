const { app, BrowserWindow } = require('electron')

app.on('ready', () => {
    console.log("ready")

    const win = new BrowserWindow({
        width: 600,
        height: 600,
        show: false
    })
    win.loadURL('https://github.com')

    win.once('ready-to-show', () => {
        win.show()
    })
})