const { app, BrowserWindow } = require(`electron`);

app.on('ready', () => {
    console.log("ready");

    const mainwindow = new BrowserWindow({
        width: 600,
        height: 600
    });

    mainwindow.loadURL(`https://github.com`)

    const secondwindow = new BrowserWindow({
        width: 300,
        height: 300,
        x: 0,
        y: 0,
        minHeight: 200,
        minWidth: 200,
        maxHeight: 500,
        maxWidth: 500,
        movable: false,
        title: "second window"
    });
    secondwindow.loadURL(`file://${__dirname}/second.html`);
});