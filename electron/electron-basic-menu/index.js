const { app, BrowserWindow, Menu } = require('electron')
const template = [{
        label: 'first',
        submenu: [{
            label: 'first-sub1',
            click: () => {
                console.log("hello?")
                app.quit()
            }
        }]
    },
    {
        label: 'second',
        submenu: [{
                label: 'second-sub1',
                click: () => {
                    console.log("hello?")
                }
            },
            { type: 'separator' },
            {
                label: "second-sub2",
                click: () => {
                    console.log("hello?")
                }
            }
        ]
    }
];


app.on('ready', () => {
    console.log("ready");

    const menu = Menu.buildFromTemplate(template);
    Menu.setApplicationMenu(menu);


    const win = new BrowserWindow();
});