const { app, BrowserWindow } = require('electron')

app.on('ready', () => {
    console.log('ready')

    const win = new BrowserWindow({
        show: false
    })

    win.loadURL(`file://${__dirname}/index.html`)

    win.on('ready-to-show', () => {
        console.log('ready-to-show')
        win.show()
    })
    win.on('close', () => {
        console.log('close')
    })
    win.on('closed', () => {
        console.log('closed')
    })
})