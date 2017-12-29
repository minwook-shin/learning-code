const { app, BrowserWindow } = require("electron");

console.log("start")

app.on('ready', () => {
    console.log('ready');
    const mainWindow = new BrowserWindow({
        width: 600,
        height: 600
    });
});

app.on('wondow-all-closed', () => {
    console.log('wondow-all-closed');
});
app.on('quit', () => {
    console.log('quit');
});