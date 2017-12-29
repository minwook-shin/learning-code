const { app, BrowserWindow } = require('electron');

app.on('ready', () => {
    console.log('ready');

    const first = new BrowserWindow({
        frame: false
    });
    first.loadURL(`file://${__dirname}/index.html`);

    const second = new BrowserWindow({
        titleBarStyle: 'hidden'
    });
    second.loadURL(`file://${__dirname}/index.html`);
});