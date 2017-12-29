const { app, BrowserWindow } = require('electron')

app.on('ready', () => {
    console.log('ready')

    const parent = new BrowserWindow()

    setTimeout(() => {
        const child = new BrowserWindow({
            width: 300,
            height: 300,
            parent: parent,
            modal: true
        });
        child.loadURL(`file://${__dirname}/index.html`)
    }, 3000)
})